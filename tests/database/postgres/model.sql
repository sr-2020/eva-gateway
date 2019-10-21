CREATE DATABASE model WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';


ALTER DATABASE model OWNER TO app;

\connect model

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
-- Name: location; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.location (
    "modelId" character varying NOT NULL,
    "timestamp" bigint NOT NULL,
    modifiers json NOT NULL,
    timers json NOT NULL,
    "manaDensity" integer NOT NULL,
    "spellTraces" json NOT NULL
);


--ALTER TABLE public.location OWNER TO appuser;

--
-- Name: qr; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public.qr (
    "modelId" character varying NOT NULL,
    "timestamp" bigint NOT NULL,
    modifiers json NOT NULL,
    timers json NOT NULL,
    "usesLeft" integer NOT NULL,
    type character varying NOT NULL,
    description character varying NOT NULL,
    "eventType" character varying NOT NULL,
    data json NOT NULL
);


--ALTER TABLE public.qr OWNER TO appuser;

--
-- Name: sr2020-character; Type: TABLE; Schema: public; Owner: appuser
--

CREATE TABLE public."sr2020-character" (
    "modelId" character varying NOT NULL,
    "timestamp" bigint NOT NULL,
    modifiers json NOT NULL,
    timers json NOT NULL,
    "maxHp" integer NOT NULL,
    "healthState" text NOT NULL,
    magic integer NOT NULL,
    "magicPowerBonus" integer NOT NULL,
    "magicAura" character varying NOT NULL,
    "spellsCasted" integer NOT NULL,
    spells json NOT NULL,
    "activeAbilities" json NOT NULL,
    "passiveAbilities" json NOT NULL,
    history json NOT NULL
);


--ALTER TABLE public."sr2020-character" OWNER TO appuser;

--
-- Data for Name: location; Type: TABLE DATA; Schema: public; Owner: appuser
--

COPY public.location ("modelId", "timestamp", modifiers, timers, "manaDensity", "spellTraces") FROM stdin;
0	1567719494736	[]	{}	40	[{"spellName": "Live long and prosper", "timestamp": 1567719494736, "casterAura": "abcdeabcdeabcdeabcde", "magicFeedback": null}]
2	1563831442887	[]	{}	431	[]
3	1563831442887	[]	{}	431	[]
1	1571475301193	[]	{}	431	[{"power":1,"spellName":"Light heal","timestamp":1567461846675,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Light heal","timestamp":1567461879567,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Trackpoint","timestamp":1567461897983,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Trackpoint","timestamp":1567461969951,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Trackpoint","timestamp":1567462514333,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Trackpoint","timestamp":1567537209672,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Field of denial","timestamp":1567589242072,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":4,"spellName":"Fireball","timestamp":1567591016726,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":2},{"power":5,"spellName":"Field of denial","timestamp":1567591026755,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":3},{"power":7,"spellName":"Live long and prosper","timestamp":1567592751645,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":4},{"power":1,"spellName":"Trackpoint","timestamp":1567593931002,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":7,"spellName":"Fireball","timestamp":1567597625333,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":4},{"power":3,"spellName":"Live long and prosper","timestamp":1567597684498,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":2},{"power":1,"spellName":"Live long and prosper","timestamp":1567597693087,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":5,"spellName":"Live long and prosper","timestamp":1567609736080,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":3},{"power":7,"spellName":"Ground heal","timestamp":1567615100644,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":4},{"power":1,"spellName":"Fireball","timestamp":1567618028088,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":6,"spellName":"Trackpoint","timestamp":1567618038122,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":3},{"power":1,"spellName":"Fireball","timestamp":1567618091593,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":3,"spellName":"Trackpoint","timestamp":1567618117667,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":2},{"power":1,"spellName":"Fireball","timestamp":1567618296869,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Fireball","timestamp":1567618299691,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Fireball","timestamp":1567618302419,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Fireball","timestamp":1567618305133,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Fireball","timestamp":1567618307582,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Fireball","timestamp":1567618314891,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Trackpoint","timestamp":1567618318109,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":7,"spellName":"Ground heal","timestamp":1567619030160,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":4},{"power":3,"spellName":"Fireball","timestamp":1567619042233,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":2},{"power":1,"spellName":"Fireball","timestamp":1567619048601,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":5,"spellName":"Ground heal","timestamp":1567702433475,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":3},{"power":1,"spellName":"Ground heal","timestamp":1567713941063,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Live long and prosper","timestamp":1567714009707,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":6,"spellName":"Live long and prosper","timestamp":1567714020030,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":3},{"power":1,"spellName":"Fireball","timestamp":1567714130102,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":7,"spellName":"Live long and prosper","timestamp":1567714238617,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":4},{"power":1,"spellName":"Trackpoint","timestamp":1567714345717,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Trackpoint","timestamp":1567714386209,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":7,"spellName":"Trackpoint","timestamp":1567715008377,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":4},{"power":7,"spellName":"Trackpoint","timestamp":1567715500108,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":4},{"power":3,"spellName":"Trackpoint","timestamp":1567715562547,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":2},{"power":1,"spellName":"Light heal","timestamp":1567715696405,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Live long and prosper","timestamp":1567719837889,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Live long and prosper","timestamp":1567721340911,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Live long and prosper","timestamp":1567721349321,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Live long and prosper","timestamp":1567723972258,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":5,"spellName":"Live long and prosper","timestamp":1567723979437,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":3},{"power":1,"spellName":"Fireball","timestamp":1567769741388,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"power":1,"spellName":"Ground heal","timestamp":1567808446748,"casterAura":"abcdeabcdeabcdeabcde","magicFeedback":1},{"spellName":"Ground heal","timestamp":1567847296413,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Ground heal","timestamp":1567848730673,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Fireball","timestamp":1568468748318,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Fireball","timestamp":1568468752246,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Live long and prosper","timestamp":1568468755921,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Ground heal","timestamp":1568468759633,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Light heal","timestamp":1568468763527,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Field of denial","timestamp":1568468769436,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Trackpoint","timestamp":1568468923564,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Fireball","timestamp":1568532560986,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Ground heal","timestamp":1568532570537,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Light heal","timestamp":1568532573277,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Live long and prosper","timestamp":1568532577576,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Trackpoint","timestamp":1568532582520,"casterAura":"abcdeabcdeabcdeabcde","power":1,"magicFeedback":1},{"spellName":"Trackpoint","timestamp":1568532597182,"casterAura":"abcdeabcdeabcdeabcde","power":7,"magicFeedback":4},{"spellName":"Trackpoint","timestamp":1568532600905,"casterAura":"abcdeabcdeabcdeabcde","power":7,"magicFeedback":4},{"spellName":"Trackpoint","timestamp":1568532698080,"casterAura":"abcdeabcdeabcdeabcde","power":3,"magicFeedback":2},{"spellName":"Field of denial","timestamp":1568535265396,"casterAura":"aaaabbbbccccddddeeee","power":1,"magicFeedback":1},{"spellName":"Live long and prosper","timestamp":1568535289165,"casterAura":"aaaabbbbccccddddeeee","power":4,"magicFeedback":2},{"spellName":"Live long and prosper","timestamp":1568535310435,"casterAura":"aaaabbbbccccddddeeee","power":1,"magicFeedback":1},{"spellName":"Light heal","timestamp":1568535381254,"casterAura":"aaaabbbbccccddddeeee","power":3,"magicFeedback":2},{"spellName":"Fireball","timestamp":1568535545073,"casterAura":"aaaabbbbccccddddeeee","power":1,"magicFeedback":1},{"spellName":"Trackpoint","timestamp":1568535568994,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1568535570348,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1568535570607,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1568535571398,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1568535571398,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1568535978633,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Fireball","timestamp":1568536029660,"casterAura":"aaaabbbbccccddddeeee","power":1,"magicFeedback":1},{"spellName":"Live long and prosper","timestamp":1568536044996,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1568536069730,"casterAura":"aaaabbbbccccddddeeee","power":1,"magicFeedback":1},{"spellName":"Fireball","timestamp":1568536952117,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1568536962046,"casterAura":"aaaabbbbccccddddeeee","power":2,"magicFeedback":1},{"spellName":"Trackpoint","timestamp":1568540623362,"casterAura":"aaaabbbbccccddddeeee","power":3,"magicFeedback":2},{"spellName":"Trackpoint","timestamp":1568551236049,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1568618448841,"casterAura":"aaaabbbbccccddddeeee","power":1,"magicFeedback":1},{"spellName":"Trackpoint","timestamp":1568618456966,"casterAura":"aaaabbbbccccddddeeee","power":1,"magicFeedback":1},{"spellName":"Trackpoint","timestamp":1568618466527,"casterAura":"aaaabbbbccccddddeeee","power":3,"magicFeedback":2},{"spellName":"Fireball","timestamp":1568738259013,"casterAura":"abcdeabcdeabcdeabcde","power":2,"magicFeedback":1},{"spellName":"Fireball","timestamp":1570724290538,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1570724303855,"casterAura":"aaaabbbbccccddddeeee","power":2,"magicFeedback":1},{"spellName":"Fireball","timestamp":1570733039880,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Trackpoint","timestamp":1570733052396,"casterAura":"aaaabbbbccccddddeeee","power":2,"magicFeedback":1},{"spellName":"Fireball","timestamp":1570747644396,"casterAura":"aaaabbbbccccddddeeee","power":5,"magicFeedback":3},{"spellName":"Light heal","timestamp":1570747653713,"casterAura":"aaaabbbbccccddddeeee","power":2,"magicFeedback":1},{"spellName":"Trackpoint","timestamp":1570747661495,"casterAura":"aaaabbbbccccddddeeee","power":1,"magicFeedback":1},{"spellName":"Ground heal","timestamp":1571475301193,"casterAura":"abcdeabcdeabcdeabcde","power":2,"magicFeedback":1}]
\.


--
-- Data for Name: qr; Type: TABLE DATA; Schema: public; Owner: appuser
--

COPY public.qr ("modelId", "timestamp", modifiers, timers, "usesLeft", type, description, "eventType", data) FROM stdin;
0	1564773697262	[]	{}	0	event	whatever	dummySpell	{}
26	1565368794384	[]	{}	1	artifact	Этот артефакт позволяет восстановить все хиты даже не будучи магом!	fullHealSpell	{}
38	1567256075518	[]	{}	0	empty		fullHealSpell	{}
45	1565307283055	[]	{}	1	artifact	Этот артефакт позволяет восстановить все хиты даже не будучи магом!	fullHealSpell	{}
\.


--
-- Data for Name: sr2020-character; Type: TABLE DATA; Schema: public; Owner: appuser
--

COPY public."sr2020-character" ("modelId", "timestamp", modifiers, timers, "maxHp", "healthState", magic, "magicPowerBonus", "magicAura", "spellsCasted", spells, "activeAbilities", "passiveAbilities", history) FROM stdin;
\.


--
-- Name: qr PK_20ce2e212c3f63fe39b80d68b57; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.qr
    ADD CONSTRAINT "PK_20ce2e212c3f63fe39b80d68b57" PRIMARY KEY ("modelId");


--
-- Name: sr2020-character PK_330ebd9982c57218c16e18cb02d; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public."sr2020-character"
    ADD CONSTRAINT "PK_330ebd9982c57218c16e18cb02d" PRIMARY KEY ("modelId");


--
-- Name: location PK_b979bcb46cb32560153f9f5f23b; Type: CONSTRAINT; Schema: public; Owner: appuser
--

ALTER TABLE ONLY public.location
    ADD CONSTRAINT "PK_b979bcb46cb32560153f9f5f23b" PRIMARY KEY ("modelId");


--
-- Name: DATABASE model; Type: ACL; Schema: -; Owner: app
--

-- GRANT ALL ON DATABASE model TO appuser;