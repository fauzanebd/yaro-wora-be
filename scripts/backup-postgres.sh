#!/bin/bash

# PostgreSQL Backup Script for Yaro Wora API
# Creates multiple backup types for maximum portability
# Usage: ./backup-postgres.sh [backup_type]

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
# Allow overriding via environment; default to conventional system path
BACKUP_DIR="${BACKUP_DIR:-/var/backups/yaro-wora-db}"

# Load project .env if present to get DB_* variables for cron and manual runs
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${SCRIPT_DIR}/.." && pwd)"
if [ -f "${PROJECT_DIR}/.env" ]; then
    set -a
    . "${PROJECT_DIR}/.env"
    set +a
fi

TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

# Require DB settings from environment/.env (no in-script defaults)
DB_NAME="${DB_NAME}"
DB_USER="${DB_USER}"
DB_HOST="${DB_HOST}"
DB_PORT="${DB_PORT}"

# Validate required variables
missing_vars=()
[ -z "${DB_NAME}" ] && missing_vars+=("DB_NAME")
[ -z "${DB_USER}" ] && missing_vars+=("DB_USER")
[ -z "${DB_HOST}" ] && missing_vars+=("DB_HOST")
[ -z "${DB_PORT}" ] && missing_vars+=("DB_PORT")

# Prefer PGPASSWORD from environment; otherwise use DB_PASSWORD if provided
if [ -n "${DB_PASSWORD}" ] && [ -z "${PGPASSWORD}" ]; then
    export PGPASSWORD="${DB_PASSWORD}"
fi

if [ -z "${PGPASSWORD}" ]; then
    missing_vars+=("PGPASSWORD or DB_PASSWORD")
fi

if [ ${#missing_vars[@]} -ne 0 ]; then
    echo "${RED}❌ Missing required environment variables:${NC} ${missing_vars[*]}" >&2
    echo "Set them in .env or environment before running this script." >&2
    exit 1
fi

# Create backup directory
mkdir -p "${BACKUP_DIR}"

# Function to log with timestamp
log() {
    echo -e "${BLUE}[$(date '+%Y-%m-%d %H:%M:%S')]${NC} $1"
}

# Function to create logical backup (SQL format)
create_sql_backup() {
    log "📦 Creating SQL backup..."
    local backup_file="${BACKUP_DIR}/yaro_wora_sql_${TIMESTAMP}.sql"
    
    pg_dump -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" \
        --format=plain --verbose --no-password \
        --file="${backup_file}"
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ SQL backup created: ${backup_file}${NC}"
        echo "File size: $(du -h "${backup_file}" | cut -f1)"
    else
        log "${RED}❌ SQL backup failed!${NC}"
        return 1
    fi
}

# Function to create custom format backup (compressed, portable)
create_custom_backup() {
    log "📦 Creating custom format backup..."
    local backup_file="${BACKUP_DIR}/yaro_wora_custom_${TIMESTAMP}.dump"
    
    pg_dump -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" \
        --format=custom --verbose --no-password \
        --file="${backup_file}"
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ Custom backup created: ${backup_file}${NC}"
        echo "File size: $(du -h "${backup_file}" | cut -f1)"
    else
        log "${RED}❌ Custom backup failed!${NC}"
        return 1
    fi
}

# Function to create directory format backup (most portable)
create_directory_backup() {
    log "📦 Creating directory format backup..."
    local backup_dir="${BACKUP_DIR}/yaro_wora_dir_${TIMESTAMP}"
    
    pg_dump -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" \
        --format=directory --verbose --no-password \
        --file="${backup_dir}"
    
    if [ $? -eq 0 ]; then
        # Create tar archive for easier handling
        tar -czf "${backup_dir}.tar.gz" -C "${BACKUP_DIR}" "yaro_wora_dir_${TIMESTAMP}"
        rm -rf "${backup_dir}"
        log "${GREEN}✅ Directory backup created: ${backup_dir}.tar.gz${NC}"
        echo "File size: $(du -h "${backup_dir}.tar.gz" | cut -f1)"
    else
        log "${RED}❌ Directory backup failed!${NC}"
        return 1
    fi
}

# Function to create schema-only backup
create_schema_backup() {
    log "📦 Creating schema-only backup..."
    local backup_file="${BACKUP_DIR}/yaro_wora_schema_${TIMESTAMP}.sql"
    
    pg_dump -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" \
        --schema-only --verbose --no-password \
        --file="${backup_file}"
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ Schema backup created: ${backup_file}${NC}"
        echo "File size: $(du -h "${backup_file}" | cut -f1)"
    else
        log "${RED}❌ Schema backup failed!${NC}"
        return 1
    fi
}

# Function to create data-only backup
create_data_backup() {
    log "📦 Creating data-only backup..."
    local backup_file="${BACKUP_DIR}/yaro_wora_data_${TIMESTAMP}.sql"
    
    pg_dump -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" \
        --data-only --verbose --no-password \
        --file="${backup_file}"
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ Data backup created: ${backup_file}${NC}"
        echo "File size: $(du -h "${backup_file}" | cut -f1)"
    else
        log "${RED}❌ Data backup failed!${NC}"
        return 1
    fi
}

# Function to create compressed archive of all backups
create_archive() {
    log "📦 Creating backup archive..."
    local archive_file="${BACKUP_DIR}/yaro_wora_full_backup_${TIMESTAMP}.tar.gz"
    
    # Create archive of all backup files from this run
    tar -czf "${archive_file}" -C "${BACKUP_DIR}" \
        yaro_wora_sql_${TIMESTAMP}.sql \
        yaro_wora_custom_${TIMESTAMP}.dump \
        yaro_wora_dir_${TIMESTAMP}.tar.gz \
        yaro_wora_schema_${TIMESTAMP}.sql \
        yaro_wora_data_${TIMESTAMP}.sql 2>/dev/null || true
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ Archive created: ${archive_file}${NC}"
        echo "Archive size: $(du -h "${archive_file}" | cut -f1)"
    else
        log "${YELLOW}⚠️  Archive creation had some issues, but individual backups are available${NC}"
    fi
}

# Function to clean old backups (keep last 7 days)
cleanup_old_backups() {
    log "🧹 Cleaning up old backups (keeping last 7 days)..."
    find "${BACKUP_DIR}" -name "yaro_wora_*" -type f -mtime +7 -delete
    log "${GREEN}✅ Old backups cleaned up${NC}"
}

# Function to verify backup integrity
verify_backup() {
    local backup_file="$1"
    local backup_type="$2"
    
    log "🔍 Verifying ${backup_type} backup..."
    
    case "${backup_type}" in
        "sql"|"schema"|"data")
            # For SQL files, check if they contain valid SQL
            if grep -q "PostgreSQL database dump" "${backup_file}"; then
                log "${GREEN}✅ ${backup_type} backup is valid${NC}"
            else
                log "${RED}❌ ${backup_type} backup appears to be corrupted${NC}"
                return 1
            fi
            ;;
        "custom")
            # For custom format, try to list contents
            if pg_restore --list "${backup_file}" > /dev/null 2>&1; then
                log "${GREEN}✅ ${backup_type} backup is valid${NC}"
            else
                log "${RED}❌ ${backup_type} backup appears to be corrupted${NC}"
                return 1
            fi
            ;;
        "directory")
            # For directory format, check if tar is valid
            if tar -tzf "${backup_file}" > /dev/null 2>&1; then
                log "${GREEN}✅ ${backup_type} backup is valid${NC}"
            else
                log "${RED}❌ ${backup_type} backup appears to be corrupted${NC}"
                return 1
            fi
            ;;
    esac
}

# Main execution
main() {
    local backup_type="${1:-all}"
    
    log "${YELLOW}🚀 Starting PostgreSQL backup process...${NC}"
    log "Database: ${DB_NAME}"
    log "User: ${DB_USER}"
    log "Host: ${DB_HOST}:${DB_PORT}"
    log "Backup directory: ${BACKUP_DIR}"
    log "Timestamp: ${TIMESTAMP}"
    echo ""
    
    # Test database connection
    log "🔌 Testing database connection..."
    if ! psql -h "${DB_HOST}" -p "${DB_PORT}" -U "${DB_USER}" -d "${DB_NAME}" -c "SELECT 1;" > /dev/null 2>&1; then
        log "${RED}❌ Cannot connect to database!${NC}"
        log "Please check your database configuration and credentials."
        exit 1
    fi
    log "${GREEN}✅ Database connection successful${NC}"
    echo ""
    
    # Create backups based on type
    case "${backup_type}" in
        "sql")
            create_sql_backup
            verify_backup "${BACKUP_DIR}/yaro_wora_sql_${TIMESTAMP}.sql" "sql"
            ;;
        "custom")
            create_custom_backup
            verify_backup "${BACKUP_DIR}/yaro_wora_custom_${TIMESTAMP}.dump" "custom"
            ;;
        "directory")
            create_directory_backup
            verify_backup "${BACKUP_DIR}/yaro_wora_dir_${TIMESTAMP}.tar.gz" "directory"
            ;;
        "schema")
            create_schema_backup
            verify_backup "${BACKUP_DIR}/yaro_wora_schema_${TIMESTAMP}.sql" "schema"
            ;;
        "data")
            create_data_backup
            verify_backup "${BACKUP_DIR}/yaro_wora_data_${TIMESTAMP}.sql" "data"
            ;;
        "all"|*)
            log "${YELLOW}📦 Creating all backup types...${NC}"
            create_sql_backup
            create_custom_backup
            create_directory_backup
            create_schema_backup
            create_data_backup
            create_archive
            ;;
    esac
    
    # Cleanup old backups
    cleanup_old_backups
    
    echo ""
    log "${GREEN}🎉 Backup process completed successfully!${NC}"
    log "${YELLOW}📋 Backup files created:${NC}"
    ls -lh "${BACKUP_DIR}"/yaro_wora_*${TIMESTAMP}* 2>/dev/null || true
    
    echo ""
    log "${BLUE}💡 Next steps:${NC}"
    log "• Test restore process: ./restore-postgres.sh"
    log "• Upload backups to cloud storage for safety"
    log "• Set up automated backups with cron"
    log "• Document your backup strategy"
}

# Show usage if help requested
if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    echo "PostgreSQL Backup Script for Yaro Wora API"
    echo ""
    echo "Usage: $0 [backup_type]"
    echo ""
    echo "Backup types:"
    echo "  all       - Create all backup types (default)"
    echo "  sql       - SQL format (human-readable)"
    echo "  custom    - Custom format (compressed, portable)"
    echo "  directory - Directory format (most portable)"
    echo "  schema    - Schema only (structure)"
    echo "  data      - Data only (content)"
    echo ""
    echo "Examples:"
    echo "  $0              # Create all backup types"
    echo "  $0 sql          # Create SQL backup only"
    echo "  $0 custom       # Create custom format backup"
    echo ""
    exit 0
fi

# Run main function
main "$@"
