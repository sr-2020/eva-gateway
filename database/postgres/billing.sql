CREATE DATABASE billing WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


ALTER DATABASE billing OWNER TO app;

\connect billing

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

SET default_with_oids = false;

--
-- Name: SystemSettings; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public."SystemSettings" (
    "Id" integer NOT NULL,
    "Key" text,
    "Value" text
);


ALTER TABLE public."SystemSettings" OWNER TO app;

--
-- Name: SystemSettings_Id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public."SystemSettings_Id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."SystemSettings_Id_seq" OWNER TO app;

--
-- Name: SystemSettings_Id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public."SystemSettings_Id_seq" OWNED BY public."SystemSettings"."Id";


--
-- Name: __EFMigrationsHistory; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public."__EFMigrationsHistory" (
    "MigrationId" character varying(150) NOT NULL,
    "ProductVersion" character varying(32) NOT NULL
);


ALTER TABLE public."__EFMigrationsHistory" OWNER TO app;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL,
    sin_from integer NOT NULL,
    sin_to integer NOT NULL,
    amount integer NOT NULL,
    comment text,
    recurrent_payment_id integer
);


ALTER TABLE public.transactions OWNER TO app;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transactions_id_seq OWNER TO app;

--
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- Name: SystemSettings Id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public."SystemSettings" ALTER COLUMN "Id" SET DEFAULT nextval('public."SystemSettings_Id_seq"'::regclass);


--
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- Data for Name: SystemSettings; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public."SystemSettings" ("Id", "Key", "Value") FROM stdin;
17	job_test:1_cron	* * * * *
18	job_test:1_start	03.10.2019 19:00:00
19	job_test:1_end	03.10.2019 20:30:00
\.


--
-- Data for Name: __EFMigrationsHistory; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public."__EFMigrationsHistory" ("MigrationId", "ProductVersion") FROM stdin;
20190930114853_Initial	2.1.11-servicing-32099
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.transactions (id, created_at, sin_from, sin_to, amount, comment, recurrent_payment_id) FROM stdin;
3	2019-05-09 00:00:00+00	0	8	99999	With great power comes great responsibility	\N
5	2019-05-27 20:08:41+00	8	5	1000	for testing	0
\.


--
-- Name: SystemSettings_Id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public."SystemSettings_Id_seq"', 20, true);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.transactions_id_seq', 316, true);


--
-- Name: SystemSettings PK_SystemSettings; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public."SystemSettings"
    ADD CONSTRAINT "PK_SystemSettings" PRIMARY KEY ("Id");


--
-- Name: __EFMigrationsHistory PK___EFMigrationsHistory; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public."__EFMigrationsHistory"
    ADD CONSTRAINT "PK___EFMigrationsHistory" PRIMARY KEY ("MigrationId");


--
-- Name: transactions idx_16423_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT idx_16423_primary PRIMARY KEY (id);


--
-- Name: DATABASE billing; Type: ACL; Schema: -; Owner: app
--

-- GRANT ALL ON DATABASE billing TO appuser;


--
-- Name: TABLE transactions; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.transactions TO appuser;


--
-- Name: SEQUENCE transactions_id_seq; Type: ACL; Schema: public; Owner: app
--

-- GRANT SELECT,USAGE ON SEQUENCE public.transactions_id_seq TO appuser;
