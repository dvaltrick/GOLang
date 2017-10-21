-- Table: public.university

-- DROP TABLE public.university;

CREATE TABLE public.university
(
    unitid character varying(10) COLLATE pg_catalog."default" NOT NULL,
    opeid character varying(10) COLLATE pg_catalog."default",
    opeid6 character varying(10) COLLATE pg_catalog."default",
    instnm character varying(200) COLLATE pg_catalog."default",
    city character varying(200) COLLATE pg_catalog."default",
    stabbr character varying(200) COLLATE pg_catalog."default",
    insturl character varying(200) COLLATE pg_catalog."default",
    CONSTRAINT university_pkey PRIMARY KEY (unitid)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.university
    OWNER to postgres;