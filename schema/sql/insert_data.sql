USE invoice_test;

INSERT INTO companies (name, representative_name, phone_number, postal_code, address)
VALUES
    ('株式会社アルファ', '山田太郎', '03-1234-5678', '100-0001', '東京都千代田区丸の内1-1-1'),
    ('株式会社ベータ', '佐藤花子', '06-2345-6789', '530-0001', '大阪府大阪市北区梅田1-2-2'),
    ('株式会社ガンマ', '鈴木一郎', '092-3456-7890', '810-0001', '福岡県福岡市中央区天神1-3-3');


INSERT INTO suppliers (company_id, name, representative_name, phone_number, postal_code, address)
VALUES
    (1, 'サプライヤー株式会社ア', '鈴木次郎', '011-1234-5678', '060-0001', '北海道札幌市中央区北1条西1-1-1'),
    (2, 'サプライヤー株式会社イ', '田中三郎', '075-2345-6789', '600-8001', '京都府京都市下京区四条通1-2-2'),
    (3, 'サプライヤー株式会社ウ', '伊藤四郎', '098-3456-7890', '900-0001', '沖縄県那覇市おもろまち1-3-3');


INSERT INTO users (company_id, name, email, password, solt)
VALUES
    (1, '山田太郎', 'taro.yamada@example.com', 'password123', 'solt123'),
    (2, '佐藤花子', 'hanako.sato@example.com', 'password456', 'solt456'),
    (3, '鈴木一郎', 'ichiro.suzuki@example.com', 'password789', 'solt789');

INSERT INTO invoices (company_id, suppliers_id, issue_date, payment_amount, fee, fee_rate, tax, tax_rate, total_amount, payment_due_date, status)
VALUES
    (1, 2, '2023-01-01', 100000.00, 1000.0, 0.01, 5000.0, 0.05, 105000, '2023-02-01', 'PENDING'),
    (1, 3, '2023-01-15', 200000.00, 2000.0, 0.01, 10000.0, 0.05, 210000, '2023-03-01', 'PROCESSING'),
    (2, 1, '2023-02-01', 150000.00, 1500.0, 0.01, 7500.0, 0.05, 157500, '2023-04-01', 'PAID');

