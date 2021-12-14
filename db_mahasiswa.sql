--
-- PostgreSQL database dump
--

-- Dumped from database version 13.5 (Ubuntu 13.5-2.pgdg20.04+1)
-- Dumped by pg_dump version 14.1 (Ubuntu 14.1-2.pgdg20.04+1)

-- Started on 2021-12-14 15:44:38 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE db_mahasiswa;
--
-- TOC entry 2999 (class 1262 OID 81384)
-- Name: db_mahasiswa; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE db_mahasiswa WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE db_mahasiswa OWNER TO postgres;

\connect db_mahasiswa

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 201 (class 1259 OID 81387)
-- Name: mst_mahasiswa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mst_mahasiswa (
    id integer NOT NULL,
    name character varying(255),
    npm integer,
    address character varying(255),
    city character varying(255),
    gander character varying(100),
    religion character varying(100),
    phone integer
);


ALTER TABLE public.mst_mahasiswa OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 81385)
-- Name: mst_mahasiswa_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mst_mahasiswa_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mst_mahasiswa_id_seq OWNER TO postgres;

--
-- TOC entry 3000 (class 0 OID 0)
-- Dependencies: 200
-- Name: mst_mahasiswa_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mst_mahasiswa_id_seq OWNED BY public.mst_mahasiswa.id;


--
-- TOC entry 2859 (class 2604 OID 81390)
-- Name: mst_mahasiswa id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mst_mahasiswa ALTER COLUMN id SET DEFAULT nextval('public.mst_mahasiswa_id_seq'::regclass);


--
-- TOC entry 2993 (class 0 OID 81387)
-- Dependencies: 201
-- Data for Name: mst_mahasiswa; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.mst_mahasiswa (id, name, npm, address, city, gander, religion, phone) VALUES (1, 'budi', 58812738, 'kuningan, jaksel', 'Jakarta', 'male', 'islam', 81237123);
INSERT INTO public.mst_mahasiswa (id, name, npm, address, city, gander, religion, phone) VALUES (2, 'tony', 8718273, 'depok, jawa barat', 'depok', 'male', 'male', 12312412);


--
-- TOC entry 3001 (class 0 OID 0)
-- Dependencies: 200
-- Name: mst_mahasiswa_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mst_mahasiswa_id_seq', 3, true);


--
-- TOC entry 2861 (class 2606 OID 81395)
-- Name: mst_mahasiswa mst_mahasiswa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mst_mahasiswa
    ADD CONSTRAINT mst_mahasiswa_pkey PRIMARY KEY (id);


-- Completed on 2021-12-14 15:44:38 WIB

--
-- PostgreSQL database dump complete
--

