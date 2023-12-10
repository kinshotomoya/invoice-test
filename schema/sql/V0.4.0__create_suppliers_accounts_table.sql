CREATE TABLE IF NOT EXISTS suppliers_accounts (
    account_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '銀行口座ID',
    suppliers_id INTEGER UNSIGNED NOT NULL COMMENT '取引先ID',
    bank_name VARCHAR(255) NOT NULL COMMENT '銀行名',
    bank_branch_name VARCHAR(255) NOT NULL COMMENT '支店名',
    account_number VARCHAR(20) NOT NULL COMMENT '口座番号',
    account_name VARCHAR(255) NOT NULL COMMENT '口座番号',
    FOREIGN KEY (suppliers_id) REFERENCES suppliers(suppliers_id),
    PRIMARY KEY (account_id)
    );