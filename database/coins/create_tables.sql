DROP TABLE IF EXISTS coins;
CREATE TABLE coins (
  id         VARCHAR(128) NOT NULL,
  name       VARCHAR(128) NOT NULL,
  symbol     VARCHAR(128) NOT NULL, 
  tracked    BOOLEAN NOT NULL,
  price      VARCHAR(255) NOT NULL,
  last_synced    VARCHAR(255) NOT NULL,

  PRIMARY KEY (`id`)
);

INSERT INTO coins
  (id, name, symbol, tracked, price, last_synced)
VALUES
  ('1','nam1', 'symbol1', true, 'price1', 'last_synced1'),
  ('2','nam2', 'symbol2', true, 'price2', 'last_synced2');