DROP TABLE IF EXISTS coins;
CREATE TABLE coins (
  id          INT AUTO_INCREMENT NOT NULL,
  coin_id     VARCHAR(128) NOT NULL,
  name        VARCHAR(128) NOT NULL,
  symbol      VARCHAR(128) NOT NULL, 
  tracked     BOOLEAN NOT NULL,
  price       VARCHAR(255) NOT NULL,
  last_synced    VARCHAR(255) NOT NULL,

  PRIMARY KEY (`id`)
);