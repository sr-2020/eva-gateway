CREATE DATABASE "position" WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


ALTER DATABASE "position" OWNER TO app;

\connect "position"

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
-- Name: beacons; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.beacons (
    id bigint NOT NULL,
    ssid character varying(255) NOT NULL,
    bssid character varying(255) NOT NULL,
    location_id bigint,
    label character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.beacons OWNER TO app;

--
-- Name: beacons_id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public.beacons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.beacons_id_seq OWNER TO app;

--
-- Name: beacons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public.beacons_id_seq OWNED BY public.beacons.id;


--
-- Name: locations; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.locations (
    id bigint NOT NULL,
    label character varying(255) NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.locations OWNER TO app;

--
-- Name: locations_id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public.locations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.locations_id_seq OWNER TO app;

--
-- Name: locations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public.locations_id_seq OWNED BY public.locations.id;


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
-- Name: paths; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.paths (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    location_id bigint DEFAULT '0'::bigint NOT NULL,
    beacon_id bigint DEFAULT '0'::bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.paths OWNER TO app;

--
-- Name: paths_id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public.paths_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.paths_id_seq OWNER TO app;

--
-- Name: paths_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public.paths_id_seq OWNED BY public.paths.id;


--
-- Name: positions; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.positions (
    id bigint NOT NULL,
    user_id bigint,
    beacons json,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.positions OWNER TO app;

--
-- Name: positions_beacons; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.positions_beacons (
    id bigint NOT NULL,
    position_id bigint NOT NULL,
    beacon_id bigint,
    bssid character varying(255) NOT NULL,
    level bigint NOT NULL
);


ALTER TABLE public.positions_beacons OWNER TO app;

--
-- Name: positions_beacons_id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public.positions_beacons_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.positions_beacons_id_seq OWNER TO app;

--
-- Name: positions_beacons_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public.positions_beacons_id_seq OWNED BY public.positions_beacons.id;


--
-- Name: positions_id_seq; Type: SEQUENCE; Schema: public; Owner: app
--

CREATE SEQUENCE public.positions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.positions_id_seq OWNER TO app;

--
-- Name: positions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: app
--

ALTER SEQUENCE public.positions_id_seq OWNED BY public.positions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: app
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    beacon_id bigint,
    location_id bigint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
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
-- Name: beacons id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.beacons ALTER COLUMN id SET DEFAULT nextval('public.beacons_id_seq'::regclass);


--
-- Name: locations id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.locations ALTER COLUMN id SET DEFAULT nextval('public.locations_id_seq'::regclass);


--
-- Name: migrations id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.migrations ALTER COLUMN id SET DEFAULT nextval('public.migrations_id_seq'::regclass);


--
-- Name: paths id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.paths ALTER COLUMN id SET DEFAULT nextval('public.paths_id_seq'::regclass);


--
-- Name: positions id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.positions ALTER COLUMN id SET DEFAULT nextval('public.positions_id_seq'::regclass);


--
-- Name: positions_beacons id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.positions_beacons ALTER COLUMN id SET DEFAULT nextval('public.positions_beacons_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: beacons; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.beacons (id, ssid, bssid, location_id, label, created_at, updated_at) FROM stdin;
1	E9:DC:0E:20:E3:DC	E9:DC:0E:20:E3:DC	1	\N	\N	\N
2	D2:7E:91:02:AB:64	D2:7E:91:02:AB:64	1	\N	\N	\N
3	F3:86:35:4C:6E:03	F3:86:35:4C:6E:03	1	\N	\N	\N
4	C0:DA:B3:09:A9:FB	C0:DA:B3:09:A9:FB	2	\N	\N	\N
5	F6:A3:B4:E1:D1:15	F6:A3:B4:E1:D1:15	2	\N	\N	\N
6	F3:8F:DE:2F:66:C9	F3:8F:DE:2F:66:C9	3	\N	\N	\N
\.


--
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.locations (id, label, created_at, updated_at) FROM stdin;
1	Танц-фойе Рим, 2 этаж	\N	\N
2	Концертный зал Москва	\N	\N
3	Левый коридор, 2 этаж	\N	\N
\.


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.migrations (id, migration, batch) FROM stdin;
7	2018_10_14_000000_create_users_table	1
8	2019_01_03_053129_create_positions_table	1
9	2019_01_30_175738_create_beacons_table	1
10	2019_02_10_132148_create_paths_table	1
11	2019_03_03_135919_create_locations_table	1
12	2019_04_23_084050_create_positions_beacons_table	1
\.


--
-- Data for Name: paths; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.paths (id, user_id, location_id, beacon_id, created_at, updated_at) FROM stdin;
1	1	1	0	2019-05-28 05:38:25+00	2019-05-28 05:40:27+00
2	1	2	0	2019-05-28 05:40:28+00	2019-05-28 05:40:28+00
3	1	1	0	2019-05-28 05:40:29+00	2019-05-28 05:40:29+00
4	1	2	0	2019-05-28 05:40:30+00	2019-05-28 05:40:30+00
5	1	1	0	2019-05-28 05:40:30+00	2019-05-28 05:40:32+00
6	1	2	0	2019-05-28 05:40:32+00	2019-05-28 05:40:33+00
7	1	1	0	2019-05-28 05:40:33+00	2019-06-09 22:18:33+00
8	1	2	0	2019-06-09 22:18:34+00	2019-06-09 22:18:35+00
9	1	1	0	2019-06-09 22:18:35+00	2019-06-09 22:18:35+00
10	1	2	0	2019-06-09 22:18:36+00	2019-06-09 22:18:36+00
11	1	1	0	2019-06-09 22:18:36+00	2019-06-09 22:18:38+00
12	1	2	0	2019-06-09 22:18:38+00	2019-06-09 22:18:39+00
13	1	1	0	2019-06-09 22:18:40+00	2019-06-09 22:37:52+00
14	1	2	0	2019-06-09 22:37:53+00	2019-06-09 22:37:54+00
15	1	1	0	2019-06-09 22:37:55+00	2019-06-09 22:37:55+00
16	1	2	0	2019-06-09 22:37:56+00	2019-06-09 22:37:56+00
17	1	1	0	2019-06-09 22:37:56+00	2019-06-09 22:37:58+00
18	1	2	0	2019-06-09 22:37:58+00	2019-06-09 22:37:59+00
19	1	1	0	2019-06-09 22:37:59+00	2019-06-10 05:41:58+00
20	1	2	0	2019-06-10 05:41:59+00	2019-06-10 05:42:00+00
21	1	1	0	2019-06-10 05:42:02+00	2019-06-10 05:42:02+00
22	1	2	0	2019-06-10 05:42:02+00	2019-06-10 05:42:02+00
23	1	1	0	2019-06-10 05:42:03+00	2019-06-10 05:42:05+00
24	1	2	0	2019-06-10 05:42:05+00	2019-06-10 05:42:06+00
25	1	1	0	2019-06-10 05:42:07+00	2019-06-10 05:58:15+00
26	1	2	0	2019-06-10 05:58:16+00	2019-06-10 05:58:17+00
27	1	1	0	2019-06-10 05:58:18+00	2019-06-10 05:58:18+00
28	1	2	0	2019-06-10 05:58:18+00	2019-06-10 05:58:18+00
29	1	1	0	2019-06-10 05:58:19+00	2019-06-10 05:58:20+00
30	1	2	0	2019-06-10 05:58:20+00	2019-06-10 05:58:21+00
31	1	1	0	2019-06-10 05:58:22+00	2019-06-10 07:18:27+00
32	1	2	0	2019-06-10 07:18:28+00	2019-06-10 07:18:29+00
33	1	1	0	2019-06-10 07:18:30+00	2019-06-10 07:18:30+00
34	1	2	0	2019-06-10 07:18:30+00	2019-06-10 07:18:30+00
35	1	1	0	2019-06-10 07:18:31+00	2019-06-10 07:18:32+00
36	1	2	0	2019-06-10 07:18:32+00	2019-06-10 07:18:33+00
37	1	1	0	2019-06-10 07:18:34+00	2019-06-10 08:00:51+00
38	1	2	0	2019-06-10 08:00:52+00	2019-06-10 08:00:53+00
39	1	1	0	2019-06-10 08:00:54+00	2019-06-10 08:00:54+00
40	1	2	0	2019-06-10 08:00:54+00	2019-06-10 08:00:54+00
41	1	1	0	2019-06-10 08:00:55+00	2019-06-10 08:00:57+00
42	1	2	0	2019-06-10 08:00:57+00	2019-06-10 08:00:58+00
43	1	1	0	2019-06-10 08:00:59+00	2019-06-20 06:41:48+00
44	1	2	0	2019-06-20 06:41:49+00	2019-06-20 06:41:50+00
45	1	1	0	2019-06-20 06:41:51+00	2019-06-20 06:41:51+00
46	1	2	0	2019-06-20 06:41:51+00	2019-06-20 06:41:51+00
47	1	1	0	2019-06-20 06:41:52+00	2019-06-20 06:41:54+00
48	1	2	0	2019-06-20 06:41:54+00	2019-06-20 06:41:54+00
49	1	1	0	2019-06-20 06:41:55+00	2019-06-20 06:51:04+00
50	1	2	0	2019-06-20 06:51:05+00	2019-06-20 06:51:06+00
51	1	1	0	2019-06-20 06:51:06+00	2019-06-20 06:51:06+00
52	1	2	0	2019-06-20 06:51:07+00	2019-06-20 06:51:07+00
53	1	1	0	2019-06-20 06:51:07+00	2019-06-20 06:51:09+00
54	1	2	0	2019-06-20 06:51:09+00	2019-06-20 06:51:10+00
55	1	1	0	2019-06-20 06:51:11+00	2019-06-20 07:01:00+00
56	1	2	0	2019-06-20 07:01:01+00	2019-06-20 07:01:02+00
57	1	1	0	2019-06-20 07:01:03+00	2019-06-20 07:01:03+00
58	1	2	0	2019-06-20 07:01:03+00	2019-06-20 07:01:03+00
59	1	1	0	2019-06-20 07:01:04+00	2019-06-20 07:01:06+00
60	1	2	0	2019-06-20 07:01:06+00	2019-06-20 07:01:07+00
61	1	1	0	2019-06-20 07:01:07+00	2019-06-20 07:01:07+00
62	1	2	0	2019-09-02 06:57:40+00	2019-09-02 06:57:40+00
\.


--
-- Data for Name: positions; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.positions (id, user_id, beacons, created_at, updated_at) FROM stdin;
1	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:38:25+00	2019-05-28 05:38:25+00
2	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:27+00	2019-05-28 05:40:27+00
3	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:28+00	2019-05-28 05:40:28+00
4	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-05-28 05:40:28+00	2019-05-28 05:40:28+00
5	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:29+00	2019-05-28 05:40:29+00
6	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:30+00	2019-05-28 05:40:30+00
7	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:30+00	2019-05-28 05:40:30+00
8	1	[]	2019-05-28 05:40:31+00	2019-05-28 05:40:31+00
9	1	[]	2019-05-28 05:40:31+00	2019-05-28 05:40:31+00
10	1	[]	2019-05-28 05:40:31+00	2019-05-28 05:40:31+00
11	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:32+00	2019-05-28 05:40:32+00
12	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:32+00	2019-05-28 05:40:32+00
13	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:32+00	2019-05-28 05:40:32+00
14	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:33+00	2019-05-28 05:40:33+00
15	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-05-28 05:40:33+00	2019-05-28 05:40:33+00
16	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-04 08:42:05+00	2019-06-04 08:42:05+00
17	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-05 08:05:25+00	2019-06-05 08:05:25+00
18	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-05 08:12:05+00	2019-06-05 08:12:05+00
19	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-05 08:18:05+00	2019-06-05 08:18:05+00
20	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-06 08:17:37+00	2019-06-06 08:17:37+00
21	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-07 07:41:37+00	2019-06-07 07:41:37+00
22	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-08 07:31:28+00	2019-06-08 07:31:28+00
23	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-08 08:24:08+00	2019-06-08 08:24:08+00
24	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-08 08:34:41+00	2019-06-08 08:34:41+00
25	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-08 09:12:31+00	2019-06-08 09:12:31+00
26	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-08 09:37:34+00	2019-06-08 09:37:34+00
27	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-08 09:47:39+00	2019-06-08 09:47:39+00
28	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-08 10:08:17+00	2019-06-08 10:08:17+00
29	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-08 10:11:46+00	2019-06-08 10:11:46+00
30	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-09 16:56:53+00	2019-06-09 16:56:53+00
31	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-09 17:12:47+00	2019-06-09 17:12:47+00
32	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:33+00	2019-06-09 22:18:33+00
33	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:34+00	2019-06-09 22:18:34+00
34	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-09 22:18:34+00	2019-06-09 22:18:35+00
35	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:35+00	2019-06-09 22:18:35+00
36	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:36+00	2019-06-09 22:18:36+00
37	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:36+00	2019-06-09 22:18:36+00
38	1	[]	2019-06-09 22:18:37+00	2019-06-09 22:18:37+00
39	1	[]	2019-06-09 22:18:37+00	2019-06-09 22:18:37+00
40	1	[]	2019-06-09 22:18:37+00	2019-06-09 22:18:37+00
41	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:38+00	2019-06-09 22:18:38+00
42	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:38+00	2019-06-09 22:18:38+00
43	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:39+00	2019-06-09 22:18:39+00
44	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:39+00	2019-06-09 22:18:39+00
45	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:18:39+00	2019-06-09 22:18:40+00
46	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:52+00	2019-06-09 22:37:52+00
47	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:53+00	2019-06-09 22:37:53+00
48	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-09 22:37:54+00	2019-06-09 22:37:54+00
49	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:55+00	2019-06-09 22:37:55+00
50	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:56+00	2019-06-09 22:37:56+00
51	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:56+00	2019-06-09 22:37:56+00
52	1	[]	2019-06-09 22:37:57+00	2019-06-09 22:37:57+00
53	1	[]	2019-06-09 22:37:57+00	2019-06-09 22:37:57+00
54	1	[]	2019-06-09 22:37:57+00	2019-06-09 22:37:57+00
55	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:58+00	2019-06-09 22:37:58+00
56	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:58+00	2019-06-09 22:37:58+00
57	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:58+00	2019-06-09 22:37:58+00
58	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:59+00	2019-06-09 22:37:59+00
59	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-09 22:37:59+00	2019-06-09 22:37:59+00
60	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-09 22:38:52+00	2019-06-09 22:38:52+00
61	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:41:58+00	2019-06-10 05:41:58+00
62	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:41:59+00	2019-06-10 05:41:59+00
63	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-10 05:42:00+00	2019-06-10 05:42:00+00
64	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:42:02+00	2019-06-10 05:42:02+00
65	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:42:02+00	2019-06-10 05:42:02+00
66	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:42:03+00	2019-06-10 05:42:03+00
67	1	[]	2019-06-10 05:42:03+00	2019-06-10 05:42:03+00
68	1	[]	2019-06-10 05:42:04+00	2019-06-10 05:42:04+00
69	1	[]	2019-06-10 05:42:04+00	2019-06-10 05:42:04+00
70	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:42:05+00	2019-06-10 05:42:05+00
71	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:42:05+00	2019-06-10 05:42:05+00
72	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:42:06+00	2019-06-10 05:42:06+00
73	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:42:06+00	2019-06-10 05:42:06+00
74	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:42:07+00	2019-06-10 05:42:07+00
75	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:15+00	2019-06-10 05:58:15+00
76	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:16+00	2019-06-10 05:58:16+00
77	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-10 05:58:17+00	2019-06-10 05:58:17+00
78	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:18+00	2019-06-10 05:58:18+00
79	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:18+00	2019-06-10 05:58:18+00
80	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:19+00	2019-06-10 05:58:19+00
81	1	[]	2019-06-10 05:58:19+00	2019-06-10 05:58:19+00
82	1	[]	2019-06-10 05:58:19+00	2019-06-10 05:58:19+00
83	1	[]	2019-06-10 05:58:20+00	2019-06-10 05:58:20+00
84	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:20+00	2019-06-10 05:58:20+00
85	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:20+00	2019-06-10 05:58:20+00
86	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:21+00	2019-06-10 05:58:21+00
87	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:21+00	2019-06-10 05:58:21+00
88	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 05:58:22+00	2019-06-10 05:58:22+00
89	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:27+00	2019-06-10 07:18:27+00
90	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:28+00	2019-06-10 07:18:28+00
91	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-10 07:18:29+00	2019-06-10 07:18:29+00
92	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:30+00	2019-06-10 07:18:30+00
93	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:30+00	2019-06-10 07:18:30+00
94	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:31+00	2019-06-10 07:18:31+00
95	1	[]	2019-06-10 07:18:31+00	2019-06-10 07:18:31+00
96	1	[]	2019-06-10 07:18:31+00	2019-06-10 07:18:31+00
97	1	[]	2019-06-10 07:18:31+00	2019-06-10 07:18:31+00
98	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:32+00	2019-06-10 07:18:32+00
99	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:32+00	2019-06-10 07:18:32+00
100	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:33+00	2019-06-10 07:18:33+00
101	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:33+00	2019-06-10 07:18:33+00
102	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 07:18:34+00	2019-06-10 07:18:34+00
103	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-10 07:59:59+00	2019-06-10 07:59:59+00
104	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:51+00	2019-06-10 08:00:51+00
105	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:52+00	2019-06-10 08:00:52+00
106	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-10 08:00:53+00	2019-06-10 08:00:53+00
107	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:54+00	2019-06-10 08:00:54+00
108	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:54+00	2019-06-10 08:00:54+00
109	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:55+00	2019-06-10 08:00:55+00
110	1	[]	2019-06-10 08:00:55+00	2019-06-10 08:00:55+00
111	1	[]	2019-06-10 08:00:56+00	2019-06-10 08:00:56+00
112	1	[]	2019-06-10 08:00:56+00	2019-06-10 08:00:56+00
113	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:57+00	2019-06-10 08:00:57+00
114	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:57+00	2019-06-10 08:00:57+00
115	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:58+00	2019-06-10 08:00:58+00
116	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:58+00	2019-06-10 08:00:58+00
117	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-10 08:00:59+00	2019-06-10 08:00:59+00
118	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-10 08:06:04+00	2019-06-10 08:06:04+00
119	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:48+00	2019-06-20 06:41:48+00
120	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:49+00	2019-06-20 06:41:49+00
121	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-20 06:41:50+00	2019-06-20 06:41:50+00
122	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:51+00	2019-06-20 06:41:51+00
123	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:51+00	2019-06-20 06:41:51+00
124	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:52+00	2019-06-20 06:41:52+00
125	1	[]	2019-06-20 06:41:52+00	2019-06-20 06:41:52+00
126	1	[]	2019-06-20 06:41:53+00	2019-06-20 06:41:53+00
127	1	[]	2019-06-20 06:41:53+00	2019-06-20 06:41:53+00
128	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:54+00	2019-06-20 06:41:54+00
129	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:54+00	2019-06-20 06:41:54+00
130	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:54+00	2019-06-20 06:41:54+00
131	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:54+00	2019-06-20 06:41:54+00
132	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:41:55+00	2019-06-20 06:41:55+00
133	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:04+00	2019-06-20 06:51:04+00
134	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:05+00	2019-06-20 06:51:05+00
135	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-20 06:51:06+00	2019-06-20 06:51:06+00
136	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:06+00	2019-06-20 06:51:06+00
137	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:07+00	2019-06-20 06:51:07+00
138	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:07+00	2019-06-20 06:51:07+00
139	1	[]	2019-06-20 06:51:08+00	2019-06-20 06:51:08+00
140	1	[]	2019-06-20 06:51:08+00	2019-06-20 06:51:08+00
141	1	[]	2019-06-20 06:51:08+00	2019-06-20 06:51:08+00
142	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:09+00	2019-06-20 06:51:09+00
143	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:09+00	2019-06-20 06:51:09+00
144	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:10+00	2019-06-20 06:51:10+00
145	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:10+00	2019-06-20 06:51:10+00
146	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 06:51:11+00	2019-06-20 06:51:11+00
147	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:00+00	2019-06-20 07:01:00+00
148	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:01+00	2019-06-20 07:01:01+00
149	1	[{"ssid": "room_a", "bssid": "e9:dc:0e:20:e3:dc", "level": -50}, {"ssid": "room_b", "bssid": "c0:da:b3:09:a9:fb", "level": -30}]	2019-06-20 07:01:02+00	2019-06-20 07:01:02+00
150	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:03+00	2019-06-20 07:01:03+00
151	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:03+00	2019-06-20 07:01:03+00
152	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:04+00	2019-06-20 07:01:04+00
153	1	[]	2019-06-20 07:01:04+00	2019-06-20 07:01:04+00
154	1	[]	2019-06-20 07:01:05+00	2019-06-20 07:01:05+00
155	1	[]	2019-06-20 07:01:05+00	2019-06-20 07:01:05+00
156	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:06+00	2019-06-20 07:01:06+00
157	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:06+00	2019-06-20 07:01:06+00
158	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:06+00	2019-06-20 07:01:06+00
159	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -50}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:07+00	2019-06-20 07:01:07+00
160	1	[{"ssid": "room_a", "bssid": "E9:DC:0E:20:E3:DC", "level": -10}, {"ssid": "room_b", "bssid": "C0:DA:B3:09:A9:FB", "level": -30}]	2019-06-20 07:01:07+00	2019-06-20 07:01:07+00
161	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-20 20:26:56+00	2019-06-20 20:26:56+00
162	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-20 20:31:42+00	2019-06-20 20:31:42+00
163	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-20 20:38:08+00	2019-06-20 20:38:08+00
164	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-20 21:51:43+00	2019-06-20 21:51:43+00
165	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-20 21:57:21+00	2019-06-20 21:57:21+00
166	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-21 07:41:40+00	2019-06-21 07:41:40+00
167	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-23 10:24:04+00	2019-06-23 10:24:04+00
168	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-23 10:55:05+00	2019-06-23 10:55:05+00
169	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-23 11:57:02+00	2019-06-23 11:57:02+00
170	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-23 12:03:23+00	2019-06-23 12:03:23+00
171	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-23 12:17:42+00	2019-06-23 12:17:42+00
172	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-23 12:24:42+00	2019-06-23 12:24:42+00
173	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-24 01:22:30+00	2019-06-24 01:22:30+00
174	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-06-25 21:19:55+00	2019-06-25 21:19:55+00
175	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-07-14 03:28:16+00	2019-07-14 03:28:16+00
176	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-07-14 03:36:32+00	2019-07-14 03:36:32+00
177	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-07-14 03:38:55+00	2019-07-14 03:38:55+00
178	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-07-14 05:19:38+00	2019-07-14 05:19:38+00
179	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-07-14 05:22:17+00	2019-07-14 05:22:17+00
180	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-07-14 05:29:59+00	2019-07-14 05:29:59+00
181	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-08-03 06:25:43+00	2019-08-03 06:25:43+00
182	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-08-03 06:49:45+00	2019-08-03 06:49:45+00
183	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-08-03 07:22:15+00	2019-08-03 07:22:15+00
184	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-08-13 19:43:39+00	2019-08-13 19:43:39+00
185	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-08-13 19:44:30+00	2019-08-13 19:44:30+00
186	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-08-13 20:06:06+00	2019-08-13 20:06:06+00
187	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-08-13 20:09:57+00	2019-08-13 20:09:57+00
188	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-08-13 20:49:59+00	2019-08-13 20:49:59+00
189	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-09-02 06:47:27+00	2019-09-02 06:47:27+00
190	1	[]	2019-09-02 06:52:15+00	2019-09-02 06:52:15+00
191	1	[]	2019-09-02 06:55:42+00	2019-09-02 06:55:42+00
192	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -40}]	2019-09-02 06:57:02+00	2019-09-02 06:57:02+00
193	1	[{"ssid": "beacon1", "bssid": "F6:A3:B4:E1:D1:15", "level": -40}]	2019-09-02 06:57:40+00	2019-09-02 06:57:40+00
194	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-09-05 04:18:58+00	2019-09-05 04:18:58+00
195	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-09-05 04:26:27+00	2019-09-05 04:26:27+00
196	1	[{"ssid": "beacon1", "bssid": "b0:0a:95:9d:00:0a", "level": -50}]	2019-09-05 04:34:45+00	2019-09-05 04:34:45+00
\.


--
-- Data for Name: positions_beacons; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.positions_beacons (id, position_id, beacon_id, bssid, level) FROM stdin;
1	1	1	E9:DC:0E:20:E3:DC	-10
2	1	4	C0:DA:B3:09:A9:FB	-30
3	2	1	E9:DC:0E:20:E3:DC	-10
4	2	4	C0:DA:B3:09:A9:FB	-30
5	3	1	E9:DC:0E:20:E3:DC	-50
6	3	4	C0:DA:B3:09:A9:FB	-30
7	4	1	E9:DC:0E:20:E3:DC	-50
8	4	4	C0:DA:B3:09:A9:FB	-30
9	5	1	E9:DC:0E:20:E3:DC	-10
10	5	4	C0:DA:B3:09:A9:FB	-30
11	6	1	E9:DC:0E:20:E3:DC	-50
12	6	4	C0:DA:B3:09:A9:FB	-30
13	7	1	E9:DC:0E:20:E3:DC	-10
14	7	4	C0:DA:B3:09:A9:FB	-30
15	11	1	E9:DC:0E:20:E3:DC	-10
16	11	4	C0:DA:B3:09:A9:FB	-30
17	12	1	E9:DC:0E:20:E3:DC	-50
18	12	4	C0:DA:B3:09:A9:FB	-30
19	13	1	E9:DC:0E:20:E3:DC	-50
20	13	4	C0:DA:B3:09:A9:FB	-30
21	14	1	E9:DC:0E:20:E3:DC	-50
22	14	4	C0:DA:B3:09:A9:FB	-30
23	15	1	E9:DC:0E:20:E3:DC	-10
24	15	4	C0:DA:B3:09:A9:FB	-30
25	16	\N	B0:0A:95:9D:00:0A	-50
26	17	\N	B0:0A:95:9D:00:0A	-50
27	18	\N	B0:0A:95:9D:00:0A	-50
28	19	\N	B0:0A:95:9D:00:0A	-50
29	20	\N	B0:0A:95:9D:00:0A	-50
30	21	\N	B0:0A:95:9D:00:0A	-50
31	22	\N	B0:0A:95:9D:00:0A	-50
32	23	\N	B0:0A:95:9D:00:0A	-50
33	24	\N	B0:0A:95:9D:00:0A	-50
34	25	\N	B0:0A:95:9D:00:0A	-50
35	26	\N	B0:0A:95:9D:00:0A	-50
36	27	\N	B0:0A:95:9D:00:0A	-50
37	28	\N	B0:0A:95:9D:00:0A	-50
38	29	\N	B0:0A:95:9D:00:0A	-50
39	30	\N	B0:0A:95:9D:00:0A	-50
40	31	\N	B0:0A:95:9D:00:0A	-50
41	32	1	E9:DC:0E:20:E3:DC	-10
42	32	4	C0:DA:B3:09:A9:FB	-30
43	33	1	E9:DC:0E:20:E3:DC	-50
44	33	4	C0:DA:B3:09:A9:FB	-30
45	34	1	E9:DC:0E:20:E3:DC	-50
46	34	4	C0:DA:B3:09:A9:FB	-30
47	35	1	E9:DC:0E:20:E3:DC	-10
48	35	4	C0:DA:B3:09:A9:FB	-30
49	36	1	E9:DC:0E:20:E3:DC	-50
50	36	4	C0:DA:B3:09:A9:FB	-30
51	37	1	E9:DC:0E:20:E3:DC	-10
52	37	4	C0:DA:B3:09:A9:FB	-30
53	41	1	E9:DC:0E:20:E3:DC	-10
54	41	4	C0:DA:B3:09:A9:FB	-30
55	42	1	E9:DC:0E:20:E3:DC	-50
56	42	4	C0:DA:B3:09:A9:FB	-30
57	43	1	E9:DC:0E:20:E3:DC	-50
58	43	4	C0:DA:B3:09:A9:FB	-30
59	44	1	E9:DC:0E:20:E3:DC	-50
60	44	4	C0:DA:B3:09:A9:FB	-30
61	45	1	E9:DC:0E:20:E3:DC	-10
62	45	4	C0:DA:B3:09:A9:FB	-30
63	46	1	E9:DC:0E:20:E3:DC	-10
64	46	4	C0:DA:B3:09:A9:FB	-30
65	47	1	E9:DC:0E:20:E3:DC	-50
66	47	4	C0:DA:B3:09:A9:FB	-30
67	48	1	E9:DC:0E:20:E3:DC	-50
68	48	4	C0:DA:B3:09:A9:FB	-30
69	49	1	E9:DC:0E:20:E3:DC	-10
70	49	4	C0:DA:B3:09:A9:FB	-30
71	50	1	E9:DC:0E:20:E3:DC	-50
72	50	4	C0:DA:B3:09:A9:FB	-30
73	51	1	E9:DC:0E:20:E3:DC	-10
74	51	4	C0:DA:B3:09:A9:FB	-30
75	55	1	E9:DC:0E:20:E3:DC	-10
76	55	4	C0:DA:B3:09:A9:FB	-30
77	56	1	E9:DC:0E:20:E3:DC	-50
78	56	4	C0:DA:B3:09:A9:FB	-30
79	57	1	E9:DC:0E:20:E3:DC	-50
80	57	4	C0:DA:B3:09:A9:FB	-30
81	58	1	E9:DC:0E:20:E3:DC	-50
82	58	4	C0:DA:B3:09:A9:FB	-30
83	59	1	E9:DC:0E:20:E3:DC	-10
84	59	4	C0:DA:B3:09:A9:FB	-30
85	60	\N	B0:0A:95:9D:00:0A	-50
86	61	1	E9:DC:0E:20:E3:DC	-10
87	61	4	C0:DA:B3:09:A9:FB	-30
88	62	1	E9:DC:0E:20:E3:DC	-50
89	62	4	C0:DA:B3:09:A9:FB	-30
90	63	1	E9:DC:0E:20:E3:DC	-50
91	63	4	C0:DA:B3:09:A9:FB	-30
92	64	1	E9:DC:0E:20:E3:DC	-10
93	64	4	C0:DA:B3:09:A9:FB	-30
94	65	1	E9:DC:0E:20:E3:DC	-50
95	65	4	C0:DA:B3:09:A9:FB	-30
96	66	1	E9:DC:0E:20:E3:DC	-10
97	66	4	C0:DA:B3:09:A9:FB	-30
98	70	1	E9:DC:0E:20:E3:DC	-10
99	70	4	C0:DA:B3:09:A9:FB	-30
100	71	1	E9:DC:0E:20:E3:DC	-50
101	71	4	C0:DA:B3:09:A9:FB	-30
102	72	1	E9:DC:0E:20:E3:DC	-50
103	72	4	C0:DA:B3:09:A9:FB	-30
104	73	1	E9:DC:0E:20:E3:DC	-50
105	73	4	C0:DA:B3:09:A9:FB	-30
106	74	1	E9:DC:0E:20:E3:DC	-10
107	74	4	C0:DA:B3:09:A9:FB	-30
108	75	1	E9:DC:0E:20:E3:DC	-10
109	75	4	C0:DA:B3:09:A9:FB	-30
110	76	1	E9:DC:0E:20:E3:DC	-50
111	76	4	C0:DA:B3:09:A9:FB	-30
112	77	1	E9:DC:0E:20:E3:DC	-50
113	77	4	C0:DA:B3:09:A9:FB	-30
114	78	1	E9:DC:0E:20:E3:DC	-10
115	78	4	C0:DA:B3:09:A9:FB	-30
116	79	1	E9:DC:0E:20:E3:DC	-50
117	79	4	C0:DA:B3:09:A9:FB	-30
118	80	1	E9:DC:0E:20:E3:DC	-10
119	80	4	C0:DA:B3:09:A9:FB	-30
120	84	1	E9:DC:0E:20:E3:DC	-10
121	84	4	C0:DA:B3:09:A9:FB	-30
122	85	1	E9:DC:0E:20:E3:DC	-50
123	85	4	C0:DA:B3:09:A9:FB	-30
124	86	1	E9:DC:0E:20:E3:DC	-50
125	86	4	C0:DA:B3:09:A9:FB	-30
126	87	1	E9:DC:0E:20:E3:DC	-50
127	87	4	C0:DA:B3:09:A9:FB	-30
128	88	1	E9:DC:0E:20:E3:DC	-10
129	88	4	C0:DA:B3:09:A9:FB	-30
130	89	1	E9:DC:0E:20:E3:DC	-10
131	89	4	C0:DA:B3:09:A9:FB	-30
132	90	1	E9:DC:0E:20:E3:DC	-50
133	90	4	C0:DA:B3:09:A9:FB	-30
134	91	1	E9:DC:0E:20:E3:DC	-50
135	91	4	C0:DA:B3:09:A9:FB	-30
136	92	1	E9:DC:0E:20:E3:DC	-10
137	92	4	C0:DA:B3:09:A9:FB	-30
138	93	1	E9:DC:0E:20:E3:DC	-50
139	93	4	C0:DA:B3:09:A9:FB	-30
140	94	1	E9:DC:0E:20:E3:DC	-10
141	94	4	C0:DA:B3:09:A9:FB	-30
142	98	1	E9:DC:0E:20:E3:DC	-10
143	98	4	C0:DA:B3:09:A9:FB	-30
144	99	1	E9:DC:0E:20:E3:DC	-50
145	99	4	C0:DA:B3:09:A9:FB	-30
146	100	1	E9:DC:0E:20:E3:DC	-50
147	100	4	C0:DA:B3:09:A9:FB	-30
148	101	1	E9:DC:0E:20:E3:DC	-50
149	101	4	C0:DA:B3:09:A9:FB	-30
150	102	1	E9:DC:0E:20:E3:DC	-10
151	102	4	C0:DA:B3:09:A9:FB	-30
152	103	\N	B0:0A:95:9D:00:0A	-50
153	104	1	E9:DC:0E:20:E3:DC	-10
154	104	4	C0:DA:B3:09:A9:FB	-30
155	105	1	E9:DC:0E:20:E3:DC	-50
156	105	4	C0:DA:B3:09:A9:FB	-30
157	106	1	E9:DC:0E:20:E3:DC	-50
158	106	4	C0:DA:B3:09:A9:FB	-30
159	107	1	E9:DC:0E:20:E3:DC	-10
160	107	4	C0:DA:B3:09:A9:FB	-30
161	108	1	E9:DC:0E:20:E3:DC	-50
162	108	4	C0:DA:B3:09:A9:FB	-30
163	109	1	E9:DC:0E:20:E3:DC	-10
164	109	4	C0:DA:B3:09:A9:FB	-30
165	113	1	E9:DC:0E:20:E3:DC	-10
166	113	4	C0:DA:B3:09:A9:FB	-30
167	114	1	E9:DC:0E:20:E3:DC	-50
168	114	4	C0:DA:B3:09:A9:FB	-30
169	115	1	E9:DC:0E:20:E3:DC	-50
170	115	4	C0:DA:B3:09:A9:FB	-30
171	116	1	E9:DC:0E:20:E3:DC	-50
172	116	4	C0:DA:B3:09:A9:FB	-30
173	117	1	E9:DC:0E:20:E3:DC	-10
174	117	4	C0:DA:B3:09:A9:FB	-30
175	118	\N	B0:0A:95:9D:00:0A	-50
176	119	1	E9:DC:0E:20:E3:DC	-10
177	119	4	C0:DA:B3:09:A9:FB	-30
178	120	1	E9:DC:0E:20:E3:DC	-50
179	120	4	C0:DA:B3:09:A9:FB	-30
180	121	1	E9:DC:0E:20:E3:DC	-50
181	121	4	C0:DA:B3:09:A9:FB	-30
182	122	1	E9:DC:0E:20:E3:DC	-10
183	122	4	C0:DA:B3:09:A9:FB	-30
184	123	1	E9:DC:0E:20:E3:DC	-50
185	123	4	C0:DA:B3:09:A9:FB	-30
186	124	1	E9:DC:0E:20:E3:DC	-10
187	124	4	C0:DA:B3:09:A9:FB	-30
188	128	1	E9:DC:0E:20:E3:DC	-10
189	128	4	C0:DA:B3:09:A9:FB	-30
190	129	1	E9:DC:0E:20:E3:DC	-50
191	129	4	C0:DA:B3:09:A9:FB	-30
192	130	1	E9:DC:0E:20:E3:DC	-50
193	130	4	C0:DA:B3:09:A9:FB	-30
194	131	1	E9:DC:0E:20:E3:DC	-50
195	131	4	C0:DA:B3:09:A9:FB	-30
196	132	1	E9:DC:0E:20:E3:DC	-10
197	132	4	C0:DA:B3:09:A9:FB	-30
198	133	1	E9:DC:0E:20:E3:DC	-10
199	133	4	C0:DA:B3:09:A9:FB	-30
200	134	1	E9:DC:0E:20:E3:DC	-50
201	134	4	C0:DA:B3:09:A9:FB	-30
202	135	1	E9:DC:0E:20:E3:DC	-50
203	135	4	C0:DA:B3:09:A9:FB	-30
204	136	1	E9:DC:0E:20:E3:DC	-10
205	136	4	C0:DA:B3:09:A9:FB	-30
206	137	1	E9:DC:0E:20:E3:DC	-50
207	137	4	C0:DA:B3:09:A9:FB	-30
208	138	1	E9:DC:0E:20:E3:DC	-10
209	138	4	C0:DA:B3:09:A9:FB	-30
210	142	1	E9:DC:0E:20:E3:DC	-10
211	142	4	C0:DA:B3:09:A9:FB	-30
212	143	1	E9:DC:0E:20:E3:DC	-50
213	143	4	C0:DA:B3:09:A9:FB	-30
214	144	1	E9:DC:0E:20:E3:DC	-50
215	144	4	C0:DA:B3:09:A9:FB	-30
216	145	1	E9:DC:0E:20:E3:DC	-50
217	145	4	C0:DA:B3:09:A9:FB	-30
218	146	1	E9:DC:0E:20:E3:DC	-10
219	146	4	C0:DA:B3:09:A9:FB	-30
220	147	1	E9:DC:0E:20:E3:DC	-10
221	147	4	C0:DA:B3:09:A9:FB	-30
222	148	1	E9:DC:0E:20:E3:DC	-50
223	148	4	C0:DA:B3:09:A9:FB	-30
224	149	1	E9:DC:0E:20:E3:DC	-50
225	149	4	C0:DA:B3:09:A9:FB	-30
226	150	1	E9:DC:0E:20:E3:DC	-10
227	150	4	C0:DA:B3:09:A9:FB	-30
228	151	1	E9:DC:0E:20:E3:DC	-50
229	151	4	C0:DA:B3:09:A9:FB	-30
230	152	1	E9:DC:0E:20:E3:DC	-10
231	152	4	C0:DA:B3:09:A9:FB	-30
232	156	1	E9:DC:0E:20:E3:DC	-10
233	156	4	C0:DA:B3:09:A9:FB	-30
234	157	1	E9:DC:0E:20:E3:DC	-50
235	157	4	C0:DA:B3:09:A9:FB	-30
236	158	1	E9:DC:0E:20:E3:DC	-50
237	158	4	C0:DA:B3:09:A9:FB	-30
238	159	1	E9:DC:0E:20:E3:DC	-50
239	159	4	C0:DA:B3:09:A9:FB	-30
240	160	1	E9:DC:0E:20:E3:DC	-10
241	160	4	C0:DA:B3:09:A9:FB	-30
242	161	\N	B0:0A:95:9D:00:0A	-50
243	162	\N	B0:0A:95:9D:00:0A	-50
244	163	\N	B0:0A:95:9D:00:0A	-50
245	164	\N	B0:0A:95:9D:00:0A	-50
246	165	\N	B0:0A:95:9D:00:0A	-50
247	166	\N	B0:0A:95:9D:00:0A	-50
248	167	\N	B0:0A:95:9D:00:0A	-50
249	168	\N	B0:0A:95:9D:00:0A	-50
250	169	\N	B0:0A:95:9D:00:0A	-50
251	170	\N	B0:0A:95:9D:00:0A	-50
252	171	\N	B0:0A:95:9D:00:0A	-50
253	172	\N	B0:0A:95:9D:00:0A	-50
254	173	\N	B0:0A:95:9D:00:0A	-50
255	174	\N	B0:0A:95:9D:00:0A	-50
256	175	\N	B0:0A:95:9D:00:0A	-50
257	176	\N	B0:0A:95:9D:00:0A	-50
258	177	\N	B0:0A:95:9D:00:0A	-50
259	178	\N	B0:0A:95:9D:00:0A	-50
260	179	\N	B0:0A:95:9D:00:0A	-50
261	180	\N	B0:0A:95:9D:00:0A	-50
262	181	\N	B0:0A:95:9D:00:0A	-50
263	182	\N	B0:0A:95:9D:00:0A	-50
264	183	\N	B0:0A:95:9D:00:0A	-50
265	184	\N	B0:0A:95:9D:00:0A	-50
266	185	\N	B0:0A:95:9D:00:0A	-50
267	186	\N	B0:0A:95:9D:00:0A	-50
268	187	\N	B0:0A:95:9D:00:0A	-50
269	188	\N	B0:0A:95:9D:00:0A	-50
270	189	\N	B0:0A:95:9D:00:0A	-50
271	192	\N	B0:0A:95:9D:00:0A	-40
272	193	5	F6:A3:B4:E1:D1:15	-40
273	194	\N	B0:0A:95:9D:00:0A	-50
274	195	\N	B0:0A:95:9D:00:0A	-50
275	196	\N	B0:0A:95:9D:00:0A	-50
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: app
--

COPY public.users (id, beacon_id, location_id, created_at, updated_at) FROM stdin;
1	\N	2	2019-05-14 18:05:48+00	2019-09-02 06:57:40+00
\.


--
-- Name: beacons_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.beacons_id_seq', 6, true);


--
-- Name: locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.locations_id_seq', 3, true);


--
-- Name: migrations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.migrations_id_seq', 12, true);


--
-- Name: paths_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.paths_id_seq', 62, true);


--
-- Name: positions_beacons_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.positions_beacons_id_seq', 275, true);


--
-- Name: positions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.positions_id_seq', 196, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: app
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: beacons idx_16470_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.beacons
    ADD CONSTRAINT idx_16470_primary PRIMARY KEY (id);


--
-- Name: locations idx_16479_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT idx_16479_primary PRIMARY KEY (id);


--
-- Name: migrations idx_16485_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT idx_16485_primary PRIMARY KEY (id);


--
-- Name: paths idx_16491_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.paths
    ADD CONSTRAINT idx_16491_primary PRIMARY KEY (id);


--
-- Name: positions idx_16499_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.positions
    ADD CONSTRAINT idx_16499_primary PRIMARY KEY (id);


--
-- Name: positions_beacons idx_16508_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.positions_beacons
    ADD CONSTRAINT idx_16508_primary PRIMARY KEY (id);


--
-- Name: users idx_16514_primary; Type: CONSTRAINT; Schema: public; Owner: app
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT idx_16514_primary PRIMARY KEY (id);


--
-- Name: idx_16470_beacons_bssid_unique; Type: INDEX; Schema: public; Owner: app
--

CREATE UNIQUE INDEX idx_16470_beacons_bssid_unique ON public.beacons USING btree (bssid);


--
-- Name: idx_16470_beacons_location_id_index; Type: INDEX; Schema: public; Owner: app
--

CREATE INDEX idx_16470_beacons_location_id_index ON public.beacons USING btree (location_id);


--
-- Name: idx_16508_positions_beacons_position_id_index; Type: INDEX; Schema: public; Owner: app
--

CREATE INDEX idx_16508_positions_beacons_position_id_index ON public.positions_beacons USING btree (position_id);


--
-- Name: DATABASE "position"; Type: ACL; Schema: -; Owner: app
--

-- GRANT ALL ON DATABASE "position" TO appuser;


--
-- Name: TABLE beacons; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.beacons TO appuser;


--
-- Name: TABLE locations; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.locations TO appuser;


--
-- Name: TABLE migrations; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.migrations TO appuser;


--
-- Name: TABLE paths; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.paths TO appuser;


--
-- Name: TABLE positions; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.positions TO appuser;


--
-- Name: TABLE positions_beacons; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.positions_beacons TO appuser;


--
-- Name: TABLE users; Type: ACL; Schema: public; Owner: app
--

-- GRANT ALL ON TABLE public.users TO appuser;