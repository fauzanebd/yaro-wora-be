#!/bin/bash

# PostgreSQL Restore Script for Yaro Wora API
# Restores from various backup formats
# Usage: ./restore-postgres.sh <backup_file> [target_database]

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
TARGET_DB="${2:-yaro_wora}"
TARGET_USER="yaro_user"
TARGET_HOST="localhost"
TARGET_PORT="5432"
BACKUP_FILE="$1"

# Function to log with timestamp
log() {
    echo -e "${BLUE}[$(date '+%Y-%m-%d %H:%M:%S')]${NC} $1"
}

# Function to show usage
show_usage() {
    echo "PostgreSQL Restore Script for Yaro Wora API"
    echo ""
    echo "Usage: $0 <backup_file> [target_database]"
    echo ""
    echo "Arguments:"
    echo "  backup_file    - Path to backup file to restore"
    echo "  target_database - Target database name (default: yaro_wora)"
    echo ""
    echo "Supported backup formats:"
    echo "  .sql          - SQL format backups"
    echo "  .dump         - Custom format backups"
    echo "  .tar.gz       - Directory format backups"
    echo ""
    echo "Examples:"
    echo "  $0 backup.sql                    # Restore SQL backup"
    echo "  $0 backup.dump                   # Restore custom backup"
    echo "  $0 backup.tar.gz                 # Restore directory backup"
    echo "  $0 backup.sql new_database      # Restore to different database"
    echo ""
    exit 1
}

# Function to detect backup format
detect_backup_format() {
    local file="$1"
    local extension="${file##*.}"
    
    case "${extension}" in
        "sql")
            echo "sql"
            ;;
        "dump")
            echo "custom"
            ;;
        "gz")
            if [[ "$file" == *.tar.gz ]]; then
                echo "directory"
            else
                echo "unknown"
            fi
            ;;
        *)
            echo "unknown"
            ;;
    esac
}

# Function to restore SQL backup
restore_sql_backup() {
    local backup_file="$1"
    local target_db="$2"
    
    log "📦 Restoring SQL backup..."
    log "Backup file: ${backup_file}"
    log "Target database: ${target_db}"
    
    # Create database if it doesn't exist
    log "🔧 Ensuring target database exists..."
    psql -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d postgres \
        -c "CREATE DATABASE ${target_db};" 2>/dev/null || log "Database may already exist"
    
    # Restore the backup
    log "🔄 Restoring data..."
    psql -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d "${target_db}" \
        -f "${backup_file}"
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ SQL backup restored successfully!${NC}"
    else
        log "${RED}❌ SQL backup restore failed!${NC}"
        return 1
    fi
}

# Function to restore custom format backup
restore_custom_backup() {
    local backup_file="$1"
    local target_db="$2"
    
    log "📦 Restoring custom format backup..."
    log "Backup file: ${backup_file}"
    log "Target database: ${target_db}"
    
    # Create database if it doesn't exist
    log "🔧 Ensuring target database exists..."
    psql -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d postgres \
        -c "CREATE DATABASE ${target_db};" 2>/dev/null || log "Database may already exist"
    
    # Restore the backup
    log "🔄 Restoring data..."
    pg_restore -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d "${target_db}" \
        --verbose --clean --no-acl --no-owner "${backup_file}"
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ Custom backup restored successfully!${NC}"
    else
        log "${RED}❌ Custom backup restore failed!${NC}"
        return 1
    fi
}

# Function to restore directory format backup
restore_directory_backup() {
    local backup_file="$1"
    local target_db="$2"
    local temp_dir="/tmp/restore_$(date +%s)"
    
    log "📦 Restoring directory format backup..."
    log "Backup file: ${backup_file}"
    log "Target database: ${target_db}"
    
    # Extract tar.gz file
    log "📂 Extracting backup archive..."
    mkdir -p "${temp_dir}"
    tar -xzf "${backup_file}" -C "${temp_dir}"
    
    # Find the extracted directory
    local extracted_dir=$(find "${temp_dir}" -type d -name "yaro_wora_dir_*" | head -1)
    
    if [ -z "${extracted_dir}" ]; then
        log "${RED}❌ Could not find extracted backup directory!${NC}"
        rm -rf "${temp_dir}"
        return 1
    fi
    
    # Create database if it doesn't exist
    log "🔧 Ensuring target database exists..."
    psql -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d postgres \
        -c "CREATE DATABASE ${target_db};" 2>/dev/null || log "Database may already exist"
    
    # Restore the backup
    log "🔄 Restoring data..."
    pg_restore -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d "${target_db}" \
        --verbose --clean --no-acl --no-owner "${extracted_dir}"
    
    # Cleanup
    rm -rf "${temp_dir}"
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ Directory backup restored successfully!${NC}"
    else
        log "${RED}❌ Directory backup restore failed!${NC}"
        return 1
    fi
}

# Function to verify restore
verify_restore() {
    local target_db="$1"
    
    log "🔍 Verifying restore..."
    
    # Check if database exists and has tables
    local table_count=$(psql -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d "${target_db}" \
        -t -c "SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public';" 2>/dev/null || echo "0")
    
    if [ "${table_count}" -gt 0 ]; then
        log "${GREEN}✅ Restore verification successful!${NC}"
        log "Found ${table_count} tables in database ${target_db}"
        
        # Show table list
        log "📋 Tables in database:"
        psql -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d "${target_db}" \
            -c "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name;" 2>/dev/null || true
    else
        log "${RED}❌ Restore verification failed!${NC}"
        log "No tables found in database ${target_db}"
        return 1
    fi
}

# Function to show backup info
show_backup_info() {
    local backup_file="$1"
    
    log "📋 Backup file information:"
    echo "File: ${backup_file}"
    echo "Size: $(du -h "${backup_file}" | cut -f1)"
    echo "Modified: $(stat -c %y "${backup_file}")"
    echo ""
    
    # Try to get more info based on format
    local format=$(detect_backup_format "${backup_file}")
    
    case "${format}" in
        "sql")
            log "📄 SQL backup detected"
            echo "First few lines:"
            head -10 "${backup_file}"
            ;;
        "custom")
            log "📦 Custom format backup detected"
            echo "Backup contents:"
            pg_restore --list "${backup_file}" 2>/dev/null || echo "Could not list contents"
            ;;
        "directory")
            log "📁 Directory format backup detected"
            echo "Archive contents:"
            tar -tzf "${backup_file}" | head -10
            ;;
    esac
}

# Main execution
main() {
    # Check arguments
    if [ -z "${BACKUP_FILE}" ]; then
        log "${RED}❌ No backup file specified!${NC}"
        show_usage
    fi
    
    if [ ! -f "${BACKUP_FILE}" ]; then
        log "${RED}❌ Backup file not found: ${BACKUP_FILE}${NC}"
        exit 1
    fi
    
    log "${YELLOW}🚀 Starting PostgreSQL restore process...${NC}"
    log "Backup file: ${BACKUP_FILE}"
    log "Target database: ${TARGET_DB}"
    log "Target host: ${TARGET_HOST}:${TARGET_PORT}"
    echo ""
    
    # Test database connection
    log "🔌 Testing database connection..."
    if ! psql -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d postgres -c "SELECT 1;" > /dev/null 2>&1; then
        log "${RED}❌ Cannot connect to database!${NC}"
        log "Please check your database configuration and credentials."
        exit 1
    fi
    log "${GREEN}✅ Database connection successful${NC}"
    echo ""
    
    # Show backup info
    show_backup_info "${BACKUP_FILE}"
    
    # Detect backup format and restore
    local format=$(detect_backup_format "${BACKUP_FILE}")
    log "🔍 Detected backup format: ${format}"
    echo ""
    
    case "${format}" in
        "sql")
            restore_sql_backup "${BACKUP_FILE}" "${TARGET_DB}"
            ;;
        "custom")
            restore_custom_backup "${BACKUP_FILE}" "${TARGET_DB}"
            ;;
        "directory")
            restore_directory_backup "${BACKUP_FILE}" "${TARGET_DB}"
            ;;
        *)
            log "${RED}❌ Unsupported backup format!${NC}"
            log "Supported formats: .sql, .dump, .tar.gz"
            exit 1
            ;;
    esac
    
    # Verify restore
    verify_restore "${TARGET_DB}"
    
    echo ""
    log "${GREEN}🎉 Restore process completed successfully!${NC}"
    log "${YELLOW}📋 Next steps:${NC}"
    log "• Test your application with the restored database"
    log "• Verify all data is correct"
    log "• Update application configuration if needed"
    log "• Consider creating a new backup after verification"
}

# Show usage if help requested
if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    show_usage
fi

# Run main function
main "$@"
