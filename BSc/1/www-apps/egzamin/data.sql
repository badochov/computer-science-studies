BEGIN;
CREATE TABLE osoba(login varchar NOT NULL PRIMARY KEY, haslo varchar NOT NULL, nauczyciel boolean NOT NULL);

INSERT INTO osoba VALUES ('nauczyciel_1', 'haslo_1', 1);
INSERT INTO osoba VALUES ('nauczyciel_2', 'haslo_2', 1);
INSERT INTO osoba VALUES ('nauczyciel_3', 'haslo_3', 1);
INSERT INTO osoba VALUES ('uczen_1', 'haslo_1', 0);
INSERT INTO osoba VALUES ('uczen_2', 'haslo_2', 0);
INSERT INTO osoba VALUES ('uczen_3', 'haslo_3', 0);

CREATE TABLE wpis(login_osoby varchar NOT NULL, timestamp datetime NOT NULL, tresc text NOT NULL);
INSERT INTO wpis VALUES ('nauczyciel_1', '2020-06-01 10:00:00', 'wpis_1_1');
INSERT INTO wpis VALUES ('nauczyciel_1', '2020-06-01 10:00:01', 'wpis_1_2');
INSERT INTO wpis VALUES ('nauczyciel_1', '2020-06-01 10:00:02', 'wpis_1_3');
INSERT INTO wpis VALUES ('nauczyciel_1', '2020-06-01 10:00:03', 'wpis_1_4');
INSERT INTO wpis VALUES ('nauczyciel_1', '2020-06-01 10:00:04', 'wpis_1_5');
INSERT INTO wpis VALUES ('nauczyciel_2', '2020-06-01 10:00:00', 'wpis_2_1');
INSERT INTO wpis VALUES ('nauczyciel_2', '2020-06-01 10:00:01', 'wpis_2_2');
INSERT INTO wpis VALUES ('nauczyciel_2', '2020-06-01 10:00:02', 'wpis_2_3');
INSERT INTO wpis VALUES ('nauczyciel_2', '2020-06-01 10:00:03', 'wpis_2_4');
INSERT INTO wpis VALUES ('nauczyciel_2', '2020-06-01 10:00:04', 'wpis_2_5');
INSERT INTO wpis VALUES ('nauczyciel_2', '2020-06-01 10:00:05', 'wpis_2_6');
INSERT INTO wpis VALUES ('nauczyciel_3', '2020-06-01 10:00:00', 'wpis_3_1');
INSERT INTO wpis VALUES ('nauczyciel_3', '2020-06-01 10:00:01', 'wpis_3_2');
INSERT INTO wpis VALUES ('nauczyciel_3', '2020-06-01 10:00:02', 'wpis_3_3');
INSERT INTO wpis VALUES ('nauczyciel_3', '2020-06-01 10:00:03', 'wpis_3_4');
INSERT INTO wpis VALUES ('nauczyciel_3', '2020-06-01 10:00:04', 'wpis_3_5');
INSERT INTO wpis VALUES ('nauczyciel_3', '2020-06-01 10:00:05', 'wpis_3_6');


CREATE TABLE sledzacy(login_osoby varchar NOT NULL, login_sledzonego varchar NOT NULL);
INSERT INTO sledzacy VALUES ('uczen_1', 'nauczyciel_1');
INSERT INTO sledzacy VALUES ('uczen_1', 'nauczyciel_2');
INSERT INTO sledzacy VALUES ('uczen_1', 'nauczyciel_3');
INSERT INTO sledzacy VALUES ('uczen_2', 'nauczyciel_1');
INSERT INTO sledzacy VALUES ('uczen_2', 'nauczyciel_2');
INSERT INTO sledzacy VALUES ('uczen_3', 'nauczyciel_2');
INSERT INTO sledzacy VALUES ('uczen_3', 'nauczyciel_3');
COMMIT;