#!/bin/bash

# Setup Automated Backup Cron Job
# This script sets up automated PostgreSQL backups

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
# Derive script dir dynamically so it works from /var/www or any path
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
BACKUP_SCRIPT="${SCRIPT_DIR}/backup-postgres.sh"
CRON_LOG="/var/log/yaro-wora-backup.log"

log() {
    echo -e "${BLUE}[$(date '+%Y-%m-%d %H:%M:%S')]${NC} $1"
}

# Function to setup cron job
setup_cron() {
    log "🕒 Setting up automated backup cron job..."
    
    # Create cron job entry (dedupe any existing matching lines)
    # Default: weekly on Sunday at 02:00. Override via BACKUP_CRON env if needed.
    local schedule="${BACKUP_CRON:-0 2 * * 0}"
    local cron_entry="${schedule} ${BACKUP_SCRIPT} all >> ${CRON_LOG} 2>&1"
    
    # Add to crontab without duplicates
    (crontab -l 2>/dev/null | grep -v "${BACKUP_SCRIPT} all"; echo "${cron_entry}") | crontab -
    
    if [ $? -eq 0 ]; then
        log "${GREEN}✅ Cron job added successfully!${NC}"
        log "Backup schedule: ${schedule}"
        log "Log file: ${CRON_LOG}"
    else
        log "${RED}❌ Failed to add cron job!${NC}"
        return 1
    fi
}

# Function to setup log rotation
setup_log_rotation() {
    log "📋 Setting up log rotation..."
    
    cat > /etc/logrotate.d/yaro-wora-backup << EOF
${CRON_LOG} {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
    create 644 ubuntu ubuntu
}
EOF
    
    log "${GREEN}✅ Log rotation configured!${NC}"
}

# Function to test backup
test_backup() {
    log "🧪 Testing backup script..."
    
    if [ -f "${BACKUP_SCRIPT}" ]; then
        # Test with a small backup
        "${BACKUP_SCRIPT}" schema
        
        if [ $? -eq 0 ]; then
            log "${GREEN}✅ Backup test successful!${NC}"
        else
            log "${RED}❌ Backup test failed!${NC}"
            return 1
        fi
    else
        log "${RED}❌ Backup script not found: ${BACKUP_SCRIPT}${NC}"
        return 1
    fi
}

# Function to show cron status
show_cron_status() {
    log "📋 Current cron jobs:"
    crontab -l 2>/dev/null | grep -E "(backup|yaro)" || echo "No backup cron jobs found"
    echo ""
    
    log "📋 Cron service status:"
    systemctl status cron --no-pager -l || true
    echo ""
}

# Function to create backup monitoring script
create_monitoring_script() {
    log "📊 Creating backup monitoring script..."
    
    cat > "${SCRIPT_DIR}/check-backup-status.sh" << 'EOF'
#!/bin/bash

# Backup Status Checker
BACKUP_DIR="${BACKUP_DIR:-/var/backups/yaro-wora-db}"
LOG_FILE="/var/log/yaro-wora-backup.log"

echo "=== Yaro Wora Backup Status ==="
echo "Date: $(date)"
echo ""

# Check if backup directory exists
if [ -d "${BACKUP_DIR}" ]; then
    echo "📁 Backup Directory: ${BACKUP_DIR}"
    echo "Total backups: $(ls -1 ${BACKUP_DIR}/yaro_wora_* 2>/dev/null | wc -l)"
    echo "Latest backup: $(ls -t ${BACKUP_DIR}/yaro_wora_* 2>/dev/null | head -1)"
    echo "Directory size: $(du -sh ${BACKUP_DIR} 2>/dev/null | cut -f1)"
    echo ""
else
    echo "❌ Backup directory not found: ${BACKUP_DIR}"
fi

# Check recent log entries
if [ -f "${LOG_FILE}" ]; then
    echo "📋 Recent backup logs:"
    tail -10 "${LOG_FILE}"
    echo ""
else
    echo "❌ Backup log file not found: ${LOG_FILE}"
fi

# Check cron job
echo "🕒 Cron job status:"
crontab -l 2>/dev/null | grep backup || echo "No backup cron job found"
EOF
    
    chmod +x "${SCRIPT_DIR}/check-backup-status.sh"
    log "${GREEN}✅ Monitoring script created!${NC}"
}

# Main execution
main() {
    log "${YELLOW}🚀 Setting up automated PostgreSQL backups...${NC}"
    echo ""
    
    # Check if running as root or with sudo
    if [ "$EUID" -ne 0 ]; then
        log "${YELLOW}⚠️  This script should be run with sudo for full functionality${NC}"
        log "Some features may not work without root privileges"
        echo ""
    fi
    
    # Test backup script first
    test_backup
    echo ""
    
    # Setup cron job
    setup_cron
    echo ""
    
    # Setup log rotation (requires sudo)
    if [ "$EUID" -eq 0 ]; then
        setup_log_rotation
        echo ""
    else
        log "${YELLOW}⚠️  Skipping log rotation setup (requires sudo)${NC}"
    fi
    
    # Create monitoring script
    create_monitoring_script
    echo ""
    
    # Show status
    show_cron_status
    
    log "${GREEN}🎉 Automated backup setup completed!${NC}"
    echo ""
    log "${YELLOW}📋 Next steps:${NC}"
    log "• Monitor backups: ${SCRIPT_DIR}/check-backup-status.sh"
    log "• Manual backup: ${BACKUP_SCRIPT}"
    log "• Check logs: tail -f ${CRON_LOG}"
    log "• Test restore: ${SCRIPT_DIR}/restore-postgres.sh"
    echo ""
    log "${BLUE}💡 Backup schedule:${NC}"
    log "• Weekly (default: Sunday 02:00). Override with BACKUP_CRON env if needed."
    log "• All backup formats created"
    log "• Old backups cleaned up automatically"
    log "• Logs rotated daily"
}

# Show usage if help requested
if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    echo "Automated Backup Setup Script"
    echo ""
    echo "This script sets up automated PostgreSQL backups with:"
    echo "• Weekly cron job (default Sunday 02:00); override with BACKUP_CRON env"
    echo "• Multiple backup formats"
    echo "• Log rotation"
    echo "• Monitoring tools"
    echo ""
    echo "Usage: $0"
    echo ""
    echo "Note: Run with sudo for full functionality"
    exit 0
fi

# Run main function
main "$@"
