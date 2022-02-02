CREATE TABLE readings (
    `id` varchar(38) PRIMARY KEY,
    `pv_generation` int(64),
    `load` int(64),
    `grid` int(64),
    `timestamp` int(64)
);