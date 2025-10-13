#!/bin/bash

# Database Migration Script: Render/Railway → VM
# Usage: ./migrate-database.sh

set -e

echo "🚀 Starting database migration..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Configuration
BACKUP_DIR="./backups"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BACKUP_FILE="yaro_wora_backup_${TIMESTAMP}"

# Source database (Render/Railway)
SOURCE_HOST="${SOURCE_DB_HOST:-your-render-host}"
SOURCE_PORT="${SOURCE_DB_PORT:-5432}"
SOURCE_DB="${SOURCE_DB_NAME:-yaro_wora}"
SOURCE_USER="${SOURCE_DB_USER:-your-username}"

# Target database (VM)
TARGET_HOST="${TARGET_DB_HOST:-localhost}"
TARGET_PORT="${TARGET_DB_PORT:-5432}"
TARGET_DB="${TARGET_DB_NAME:-yaro_wora}"
TARGET_USER="${TARGET_DB_USER:-yaro_user}"

echo -e "${YELLOW}📋 Migration Configuration:${NC}"
echo "Source: ${SOURCE_HOST}:${SOURCE_PORT}/${SOURCE_DB}"
echo "Target: ${TARGET_HOST}:${SOURCE_PORT}/${TARGET_DB}"
echo "Backup file: ${BACKUP_FILE}"
echo ""

# Create backup directory
mkdir -p "${BACKUP_DIR}"

echo -e "${YELLOW}📦 Step 1: Creating backup from source database...${NC}"
pg_dump -h "${SOURCE_HOST}" -p "${SOURCE_PORT}" -U "${SOURCE_USER}" -d "${SOURCE_DB}" \
  --format=custom --verbose --file="${BACKUP_DIR}/${BACKUP_FILE}.dump" \
  --no-password

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Backup created successfully: ${BACKUP_DIR}/${BACKUP_FILE}.dump${NC}"
else
    echo -e "${RED}❌ Backup failed!${NC}"
    exit 1
fi

echo -e "${YELLOW}📊 Step 2: Checking backup file...${NC}"
ls -lh "${BACKUP_DIR}/${BACKUP_FILE}.dump"

echo -e "${YELLOW}🔄 Step 3: Restoring to target database...${NC}"
pg_restore --verbose --clean --no-acl --no-owner \
  -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d "${TARGET_DB}" \
  "${BACKUP_DIR}/${BACKUP_FILE}.dump"

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Database restored successfully!${NC}"
else
    echo -e "${RED}❌ Restore failed!${NC}"
    exit 1
fi

echo -e "${YELLOW}🔍 Step 4: Verifying migration...${NC}"

# Get table counts from both databases
echo "Checking table counts..."
SOURCE_TABLES=$(psql -h "${SOURCE_HOST}" -p "${SOURCE_PORT}" -U "${SOURCE_USER}" -d "${SOURCE_DB}" -t -c "SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public';" 2>/dev/null || echo "0")
TARGET_TABLES=$(psql -h "${TARGET_HOST}" -p "${TARGET_PORT}" -U "${TARGET_USER}" -d "${TARGET_DB}" -t -c "SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public';" 2>/dev/null || echo "0")

echo "Source tables: ${SOURCE_TABLES}"
echo "Target tables: ${TARGET_TABLES}"

if [ "${SOURCE_TABLES}" = "${TARGET_TABLES}" ]; then
    echo -e "${GREEN}✅ Table counts match!${NC}"
else
    echo -e "${RED}❌ Table counts don't match!${NC}"
    exit 1
fi

echo -e "${GREEN}🎉 Migration completed successfully!${NC}"
echo -e "${YELLOW}📝 Next steps:${NC}"
echo "1. Update your application's database connection string"
echo "2. Test your application thoroughly"
echo "3. Update DNS/domain settings if needed"
echo "4. Keep the backup file: ${BACKUP_DIR}/${BACKUP_FILE}.dump"
