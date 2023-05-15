-- migrate:up
-- Criação da tabela matches
CREATE TABLE matches (
                         id VARCHAR(36) NOT NULL PRIMARY KEY,
                         match_date DATETIME,
                         team_a_id VARCHAR(36),
                         team_a_name VARCHAR(255),
                         team_b_id VARCHAR(36),
                         team_b_name VARCHAR(255),
                         result VARCHAR(255)
);

-- Criação da tabela players
CREATE TABLE players (
                         id VARCHAR(36) NOT NULL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         price DECIMAL(10,2) NOT NULL
);

-- Criação da tabela teams
CREATE TABLE teams (
                       id VARCHAR(36) NOT NULL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL
);

-- Criação da tabela team_players
CREATE TABLE team_players (
                              team_id VARCHAR(36) NOT NULL,
                              player_id VARCHAR(36) NOT NULL
);

-- Criação da tabela my_team
CREATE TABLE my_team (
                         id VARCHAR(36) NOT NULL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         score DECIMAL(10,2) NOT NULL
);

-- Criação da tabela my_team_players
CREATE TABLE my_team_players (
                                 my_team_id VARCHAR(36) NOT NULL,
                                 player_id VARCHAR(36) NOT NULL
);

-- Criação da tabela actions
CREATE TABLE actions (
                         id VARCHAR(36) NOT NULL PRIMARY KEY,
                         match_id VARCHAR(36) NOT NULL,
                         team_id VARCHAR(36) NOT NULL,
                         player_id VARCHAR(36) NOT NULL,
                         action VARCHAR(255) NOT NULL,
                         minute INTEGER NOT NULL,
                         score DECIMAL(10,2) NOT NULL
);

-- migrate:down
-- Reversão da criação das tabelas
DROP TABLE IF EXISTS actions;
DROP TABLE IF EXISTS my_team_players;
DROP TABLE IF EXISTS my_team;
DROP TABLE IF EXISTS team_players;
DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS matches;