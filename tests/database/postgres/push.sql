CREATE DATABASE push WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


ALTER DATABASE push OWNER TO app;

\connect push

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
-- Name: firebase_tokens; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.firebase_tokens (
    id bigint NOT NULL,
    token text
);


ALTER TABLE public.firebase_tokens OWNER TO app;

--
-- Data for Name: firebase_tokens; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.firebase_tokens (id, token) FROM stdin;
120	testtoken3863819
\.


--
-- Name: firebase_tokens idx_16540_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.firebase_tokens
    ADD CONSTRAINT idx_16540_primary PRIMARY KEY (id);


--
-- Name: DATABASE push; Type: ACL; Schema: -; Owner: app
--

-- GRANT ALL ON DATABASE push TO appuser;


--
-- Name: TABLE firebase_tokens; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.firebase_tokens TO appuser;


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

