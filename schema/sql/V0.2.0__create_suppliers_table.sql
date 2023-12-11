CREATE TABLE IF NOT EXISTS suppliers (
    suppliers_id INTEGER UNSIGNED AUTO_INCREMENT COMMENT '取引先ID',
    company_id INTEGER UNSIGNED COMMENT '法人ID',
    name VARCHAR(255) NOT NULL COMMENT '法人名',
    representative_name VARCHAR(255) COMMENT '代表者名',
    phone_number VARCHAR(20) COMMENT '電話番号',
    postal_code VARCHAR(10) COMMENT '郵便番号',
    address VARCHAR(255) COMMENT '郵便番号',
    FOREIGN KEY (company_id) REFERENCES companies(company_id),
    PRIMARY KEY (suppliers_id)
    );

