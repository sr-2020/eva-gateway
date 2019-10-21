CREATE DATABASE auth WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


ALTER DATABASE auth OWNER TO app;

\connect auth

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
-- Name: migrations; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.migrations (
    id bigint NOT NULL,
    migration character varying(255) NOT NULL,
    batch bigint NOT NULL
);


ALTER TABLE public.migrations OWNER TO app;

--
-- Name: migrations_id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public.migrations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.migrations_id_seq OWNER TO app;

--
-- Name: migrations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public.migrations_id_seq OWNED BY public.migrations.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    admin boolean DEFAULT false NOT NULL,
    beacon_id bigint,
    location_id bigint,
    status character varying(255) DEFAULT ''::character varying NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) DEFAULT 'test'::character varying NOT NULL,
    api_key character varying(255),
    amount bigint DEFAULT '0'::bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    options json
);


ALTER TABLE public.users OWNER TO app;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO app;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id SET DEFAULT nextval('public.migrations_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.migrations (id, migration, batch) FROM stdin;
1	2018_10_14_000000_create_users_table	1
36	2019_06_09_134106_add_options_column_to_users	2
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.users (id, admin, beacon_id, location_id, status, name, email, password, api_key, amount, created_at, updated_at, options) FROM stdin;
1	t	\N	\N	newdb	Мистер X	admin@evarun.ru	$2y$10$TRxK0F1twgCAbGygKTX1E.2wk0KNT3fPcfASlLtdFvU2AU8XsxZfG	TkRVem4yTERSQTNQRHFxcmo4SUozNWZp	223	2019-03-30 04:57:38+00	2019-06-09 20:40:42+00	{"test": 6}
\.


--
-- Name: migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.migrations_id_seq', 36, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.users_id_seq', 143, true);


--
-- Name: migrations idx_16561_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT idx_16561_primary PRIMARY KEY (id);


--
-- Name: users idx_16567_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT idx_16567_primary PRIMARY KEY (id);


--
-- Name: idx_16567_users_email_unique; Type: INDEX; Schema: public; Owner: app
--

CREATE UNIQUE INDEX idx_16567_users_email_unique ON public.users USING btree (email);


--
-- Name: idx_16567_users_name_unique; Type: INDEX; Schema: public; Owner: app
--

CREATE UNIQUE INDEX idx_16567_users_name_unique ON public.users USING btree (name);


--
-- Name: DATABASE auth; Type: ACL; Schema: -; Owner: app
--

-- GRANT ALL ON DATABASE auth TO appuser;


--
-- Name: TABLE migrations; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.migrations TO appuser;


--
-- Name: TABLE users; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.users TO appuser;
