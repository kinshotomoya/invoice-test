CREATE TABLE IF NOT EXISTS companies (
    company_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '企業ID',
    name VARCHAR(255) NOT NULL COMMENT '法人名',
    representative_name VARCHAR(255) COMMENT '代表者名',
    phone_number VARCHAR(20) COMMENT '電話番号',
    postal_code VARCHAR(10) COMMENT '郵便番号',
    address VARCHAR(255) COMMENT '郵便番号',
    PRIMARY KEY (company_id));