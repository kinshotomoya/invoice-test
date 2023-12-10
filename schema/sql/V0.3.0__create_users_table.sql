CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER UNSIGNED AUTO_INCREMENT COMMENT 'ユーザーID',
    company_id INTEGER UNSIGNED COMMENT '法人ID',
    name VARCHAR(255) NOT NULL COMMENT '氏名',
    email VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    Password VARCHAR(255) NOT NULL COMMENT 'パスワード',
    FOREIGN KEY (company_id) REFERENCES companies(company_id),
    PRIMARY KEY (user_id)
    );