create table assets
(
    id               serial not null primary key,
    "time"           bigint,
    high             numeric(19, 6),
    Low              numeric(19, 6),
    "open"           numeric(19, 6),
    Volumefrom       numeric(19, 6),
    Volumeto         numeric(19, 6),
    "close"          numeric(19, 6),
    ConversionType   varchar(1024),
    ConversionSymbol varchar(1024)
);