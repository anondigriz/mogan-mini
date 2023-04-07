-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "KnowledgeBase" (
    "ID" NVARCHAR(40) NOT NULL PRIMARY KEY,
    "ShortName" NVARCHAR(50) NOT NULL,
    "CreatedDate" INTEGER NOT NULL,
    "ModifiedDate" INTEGER NOT NULL,
    "RemoteUUID" NVARCHAR(40),
    "ExtraData" TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS "Group" (
    "ID" NVARCHAR(40) NOT NULL PRIMARY KEY,
    "ShortName" NVARCHAR(50) NOT NULL,
    "CreatedDate" INTEGER NOT NULL,
    "ModifiedDate" INTEGER NOT NULL,
    "ExtraData" TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS "TypeParameter" (
    "ID" INTEGER NOT NULL PRIMARY KEY,
    "ShortName" NVARCHAR(20) NOT NULL
);
INSERT INTO "TypeParameter"("ID", "ShortName")
VALUES (0, 'String'),
    (1, 'Double'),
    (2, 'Boolean'),
    (3, 'BigInteger');
CREATE TABLE IF NOT EXISTS "Parameter" (
    "ID" NVARCHAR(40) NOT NULL PRIMARY KEY,
    "ShortName" NVARCHAR(50) NOT NULL,
    "CreatedDate" INTEGER NOT NULL,
    "ModifiedDate" INTEGER NOT NULL,
    "GroupID" NVARCHAR(40) NOT NULL,
    "TypeID" INTEGER NOT NULL,
    "ExtraData" TEXT NOT NULL,
    CONSTRAINT "Parameter_GroupID_fkey" FOREIGN KEY ("GroupID") REFERENCES "Group" ("ID") MATCH SIMPLE ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT "Parameter_TypeID_fkey" FOREIGN KEY ("TypeID") REFERENCES "TypeParameter" ("ID") MATCH SIMPLE ON UPDATE CASCADE ON DELETE RESTRICT
);
CREATE TABLE IF NOT EXISTS "TypePattern" (
    "ID" INTEGER NOT NULL PRIMARY KEY,
    "ShortName" NVARCHAR(20) NOT NULL
);
INSERT INTO "TypePattern"("ID", "ShortName")
VALUES (0, 'Program'),
    (1, 'Constraint'),
    (2, 'Formula'),
    (3, 'IfThenElse');
CREATE TABLE IF NOT EXISTS "Pattern" (
    "ID" NVARCHAR(40) NOT NULL PRIMARY KEY,
    "ShortName" NVARCHAR(50) NOT NULL,
    "CreatedDate" INTEGER NOT NULL,
    "ModifiedDate" INTEGER NOT NULL,
    "TypeID" INTEGER NOT NULL,
    "ExtraData" TEXT NOT NULL,
    CONSTRAINT "Pattern_TypeID_fkey" FOREIGN KEY ("TypeID") REFERENCES "TypePattern" ("ID") MATCH SIMPLE ON UPDATE CASCADE ON DELETE RESTRICT
);
CREATE TABLE IF NOT EXISTS "Rule" (
    "ID" NVARCHAR(40) NOT NULL PRIMARY KEY,
    "ShortName" NVARCHAR(50) NOT NULL,
    "CreatedDate" INTEGER NOT NULL,
    "ModifiedDate" INTEGER NOT NULL,
    "PatternID" uuid NOT NULL,
    "ExtraData" TEXT NOT NULL,
    CONSTRAINT "Rule_PatternID_fkey" FOREIGN KEY ("PatternID") REFERENCES "Pattern" ("ID") MATCH SIMPLE ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Parameter";
DROP TABLE IF EXISTS "TypeParameter";
DROP TABLE IF EXISTS "Rule";
DROP TABLE IF EXISTS "Pattern";
DROP TABLE IF EXISTS "TypePattern";
DROP TABLE IF EXISTS "Group";
DROP TABLE IF EXISTS "KnowledgeBase";
-- +goose StatementEnd