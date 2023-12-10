CREATE TABLE IF NOT EXISTS invoices (
    invoice_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT  COMMENT '請求データID',
    company_id INTEGER UNSIGNED NOT NULL COMMENT '法人ID',
    suppliers_id INTEGER UNSIGNED NOT NULL COMMENT '取引先ID',
    issue_date DATE NOT NULL COMMENT '発行日',
    payment_amount DECIMAL(65,2) NOT NULL COMMENT '支払金額',
    fee DECIMAL(65,1) NOT NULL COMMENT '手数料',
--  NOTE: それぞれ桁数を最大値まで確保している
    fee_rate DECIMAL(31,30) NOT NULL COMMENT '手数料率',
    tax DECIMAL(65,1) NOT NULL COMMENT '消費税',
    tax_rate DECIMAL(31,30) NOT NULL COMMENT '消費税率',
    total_amount DECIMAL(65,0) NOT NULL COMMENT '請求金額',
    payment_due_date DATE NOT NULL COMMENT '支払期日',
    status ENUM('PENDING', 'PROCESSING', 'PAID', 'ERROR') NOT NULL COMMENT 'ステータス',
    PRIMARY KEY (invoice_id),
    FOREIGN KEY (company_id) REFERENCES companies(company_id),
    FOREIGN KEY (suppliers_id) REFERENCES suppliers(suppliers_id)
    );