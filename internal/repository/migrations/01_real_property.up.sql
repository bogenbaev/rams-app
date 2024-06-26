CREATE TABLE IF NOT EXISTS real_property (
    id SERIAL PRIMARY KEY,
    property_type_id INTEGER NOT NULL,
    property_type VARCHAR(50) NOT NULL,
    address VARCHAR(100) NOT NULL,
    price INTEGER NOT NULL,
    rooms INTEGER DEFAULT 0,
    area FLOAT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_real_property_property_type_id ON real_property (property_type_id);

INSERT INTO real_property (property_type_id, property_type, address, price, rooms, area, description) VALUES
    (1, 'Квартира', 'Ул. Ленина, д. 5, кв. 3', 1200000, 1, 45.5, 'Студия в новом доме с ремонтом и мебелью'),
    (2, 'Дом', 'Ул. Центральная, д. 12', 2800000, 3, 150, 'Просторный дом с двором и гаражом'),
    (3, 'Земельный участок', 'СНТ "Рассвет", участок 25', 400000, NULL, 800, 'Участок для строительства загородного дома'),
    (1, 'Квартира', 'Проспект Победы, д. 20, кв. 8', 1800000, 2, 55, 'Уютная квартира с отличным видом на город'),
    (2, 'Дом', 'Ул. Садовая, д. 7', 3500000, 4, 180, 'Дом с садом и бассейном, идеально для семьи'),
    (3, 'Земельный участок', 'Деревня Солнечная, участок 10', 250000, NULL, 600, 'Участок для дачного строительства'),
    (1, 'Квартира', 'Проспект Мира, д. 15, кв. 12', 1600000, 3, 75, 'Трехкомнатная квартира с балконом'),
    (2, 'Дом', 'Ул. Новая, д. 3', 4200000, 5, 220, 'Просторный дом с садом и беседкой'),
    (3, 'Земельный участок', 'Деревня Зеленая, участок 5', 300000, NULL, 1000, 'Участок с лесом, идеально для строительства загородного дома'),
    (1, 'Квартира', 'Ул. Юбилейная, д. 30, кв. 22', 1400000, 2, 60, 'Квартира с отличным ремонтом в центре города'),
    (2, 'Дом', 'Ул. Лесная, д. 14', 3200000, 4, 210, 'Красивый дом в окружении природы'),
    (1, 'Квартира', 'Проспект Свободы, д. 7, кв. 15', 1750000, 3, 70, 'Светлая и просторная квартира в новом доме'),
    (2, 'Дом', 'Ул. Солнечная, д. 25', 3800000, 6, 280, 'Дом с большим двором и бассейном'),
    (3, 'Земельный участок', 'Дачное товарищество "Луч", участок 18', 220000, NULL, 800, 'Участок с видом на озеро'),
    (1, 'Квартира', 'Ул. Звездная, д. 8, кв. 42', 1350000, 2, 65, 'Уютная квартира в тихом районе'),
    (2, 'Дом', 'Пер. Сиреневый, д. 3', 4800000, 7, 320, 'Роскошный дом с собственным парком'),
    (3, 'Земельный участок', 'Деревня Радость, участок 7', 280000, NULL, 900, 'Участок для строительства загородного дома'),
    (1, 'Квартира', 'Ул. Цветочная, д. 14, кв. 30', 1550000, 3, 80, 'Квартира с красивым видом на парк'),
    (2, 'Дом', 'Ул. Луговая, д. 9', 3700000, 5, 250, 'Современный дом с террасой и бассейном');