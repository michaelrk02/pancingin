-- DROP SCHEMA dbo;

-- CREATE SCHEMA dbo;
-- DBMSProject.dbo.IKAN definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.IKAN;

CREATE TABLE IKAN (
	KdIkan int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	HargaJual int NULL,
	CONSTRAINT PK_IKAN PRIMARY KEY (KdIkan)
);


-- DBMSProject.dbo.LAUT definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.LAUT;

CREATE TABLE LAUT (
	KdLaut int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	Lokasi varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	CONSTRAINT PK_LAUT PRIMARY KEY (KdLaut)
);


-- DBMSProject.dbo.NELAYAN definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.NELAYAN;

CREATE TABLE NELAYAN (
	KdNelayan int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	NoTelp varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	CONSTRAINT PK_NELAYAN PRIMARY KEY (KdNelayan)
);


-- DBMSProject.dbo.PABRIK definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PABRIK;

CREATE TABLE PABRIK (
	KdPabrik int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	Lokasi varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	NoTelp varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	CONSTRAINT PK_PABRIK PRIMARY KEY (KdPabrik)
);


-- DBMSProject.dbo.PELANGGAN definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PELANGGAN;

CREATE TABLE PELANGGAN (
	KdPelanggan int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	Lokasi varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	CONSTRAINT PK_PELANGGAN PRIMARY KEY (KdPelanggan)
);


-- DBMSProject.dbo.PETANI definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PETANI;

CREATE TABLE PETANI (
	KdPetani int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	NoTelp varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	CONSTRAINT PK_PETANI PRIMARY KEY (KdPetani)
);


-- DBMSProject.dbo.SUPPLIER definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.SUPPLIER;

CREATE TABLE SUPPLIER (
	KdSupplier int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	NoTelp varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	Lokasi varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	CONSTRAINT PK_SUPPLIER PRIMARY KEY (KdSupplier)
);


-- DBMSProject.dbo.KAPAL definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.KAPAL;

CREATE TABLE KAPAL (
	KdKapal int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	Jenis varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	Kapasitas int NULL,
	KdLaut int NULL,
	CONSTRAINT PK_KAPAL PRIMARY KEY (KdKapal),
	CONSTRAINT FK__KAPAL__KdLaut FOREIGN KEY (KdLaut) REFERENCES LAUT(KdLaut)
);


-- DBMSProject.dbo.LAUT_MEMILIKI_IKAN definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.LAUT_MEMILIKI_IKAN;

CREATE TABLE LAUT_MEMILIKI_IKAN (
	KdLaut int NOT NULL,
	KdIkan int NOT NULL,
	CONSTRAINT PK_LAUT_MEMILIKI_IKAN PRIMARY KEY (KdLaut,KdIkan),
	CONSTRAINT FK__LAUT_MEMILIKI_IKAN__KdIkan FOREIGN KEY (KdIkan) REFERENCES IKAN(KdIkan),
	CONSTRAINT FK__LAUT_MEMILIKI_IKAN__KdLaut FOREIGN KEY (KdLaut) REFERENCES LAUT(KdLaut)
);


-- DBMSProject.dbo.NELAYAN_MENANGKAP_IKAN definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.NELAYAN_MENANGKAP_IKAN;

CREATE TABLE NELAYAN_MENANGKAP_IKAN (
	KdNelayan int NOT NULL,
	KdIkan int NOT NULL,
	Keuntungan int NULL,
	Persediaan int NULL,
	CONSTRAINT PK_NELAYAN_MENANGKAP_IKAN PRIMARY KEY (KdNelayan,KdIkan),
	CONSTRAINT FK__NELAYAN_MENANGKAP_IKAN__KdIkan FOREIGN KEY (KdIkan) REFERENCES IKAN(KdIkan),
	CONSTRAINT FK__NELAYAN_MENANGKAP_IKAN__KdNelayan FOREIGN KEY (KdNelayan) REFERENCES NELAYAN(KdNelayan)
);


-- DBMSProject.dbo.NELAYAN_MENGGUNAKAN_KAPAL definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.NELAYAN_MENGGUNAKAN_KAPAL;

CREATE TABLE NELAYAN_MENGGUNAKAN_KAPAL (
	KdNelayan int NOT NULL,
	KdKapal int NOT NULL,
	Tanggal date NOT NULL,
	CONSTRAINT PK_NELAYAN_MENGGUNAKAN_KAPAL PRIMARY KEY (KdNelayan,KdKapal,Tanggal),
	CONSTRAINT FK__NELAYAN_MENGGUNAKAN_KAPAL__KdKapal FOREIGN KEY (KdKapal) REFERENCES KAPAL(KdKapal),
	CONSTRAINT FK__NELAYAN_MENGGUNAKAN_KAPAL__KdNelayan FOREIGN KEY (KdNelayan) REFERENCES NELAYAN(KdNelayan)
);


-- DBMSProject.dbo.PABRIK_MEMBELI_IKAN_DARI_NELAYAN definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PABRIK_MEMBELI_IKAN_DARI_NELAYAN;

CREATE TABLE PABRIK_MEMBELI_IKAN_DARI_NELAYAN (
	KdPabrik int NOT NULL,
	KdIkan int NOT NULL,
	KdNelayan int NOT NULL,
	Jumlah int NULL,
	Tanggal date NOT NULL,
	CONSTRAINT PK_PABRIK_MEMBELI_IKAN_DARI_NELAYAN PRIMARY KEY (KdPabrik,KdIkan,KdNelayan,Tanggal),
	CONSTRAINT FK__PABRIK_MEMBELI_IKAN_DARI_NELAYAN__KdIkan FOREIGN KEY (KdIkan) REFERENCES IKAN(KdIkan),
	CONSTRAINT FK__PABRIK_MEMBELI_IKAN_DARI_NELAYAN__KdNelayan FOREIGN KEY (KdNelayan) REFERENCES NELAYAN(KdNelayan),
	CONSTRAINT FK__PABRIK_MEMBELI_IKAN_DARI_NELAYAN__KdPabrik FOREIGN KEY (KdPabrik) REFERENCES PABRIK(KdPabrik)
);


-- DBMSProject.dbo.PABRIK_MEMBELI_IKAN_DARI_PETANI definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PABRIK_MEMBELI_IKAN_DARI_PETANI;

CREATE TABLE PABRIK_MEMBELI_IKAN_DARI_PETANI (
	KdPabrik int NOT NULL,
	KdIkan int NOT NULL,
	KdPetani int NOT NULL,
	Jumlah int NULL,
	Tanggal date NOT NULL,
	CONSTRAINT PK_PABRIK_MEMBELI_IKAN_DARI_PETANI PRIMARY KEY (KdPabrik,KdIkan,KdPetani,Tanggal),
	CONSTRAINT FK__PABRIK_MEMBELI_IKAN_DARI_PETANI__KdIkan FOREIGN KEY (KdIkan) REFERENCES IKAN(KdIkan),
	CONSTRAINT FK__PABRIK_MEMBELI_IKAN_DARI_PETANI__KdPabrik FOREIGN KEY (KdPabrik) REFERENCES PABRIK(KdPabrik),
	CONSTRAINT FK__PABRIK_MEMBELI_IKAN_DARI_PETANI__KdPetani FOREIGN KEY (KdPetani) REFERENCES PETANI(KdPetani)
);


-- DBMSProject.dbo.PETANI_MEMBELI_IKAN_DARI_NELAYAN definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PETANI_MEMBELI_IKAN_DARI_NELAYAN;

CREATE TABLE PETANI_MEMBELI_IKAN_DARI_NELAYAN (
	KdPetani int NOT NULL,
	KdIkan int NOT NULL,
	KdNelayan int NOT NULL,
	Jumlah int NULL,
	Tanggal date NOT NULL,
	CONSTRAINT PK_PETANI_MEMBELI_IKAN_DARI_NELAYAN PRIMARY KEY (KdPetani,KdIkan,KdNelayan,Tanggal),
	CONSTRAINT FK__PETANI_MEMBELI_IKAN_DARI_NELAYAN__KdIkan FOREIGN KEY (KdIkan) REFERENCES IKAN(KdIkan),
	CONSTRAINT FK__PETANI_MEMBELI_IKAN_DARI_NELAYAN__KdNelayan FOREIGN KEY (KdNelayan) REFERENCES NELAYAN(KdNelayan),
	CONSTRAINT FK__PETANI_MEMBELI_IKAN_DARI_NELAYAN__KdPetani FOREIGN KEY (KdPetani) REFERENCES PETANI(KdPetani)
);


-- DBMSProject.dbo.PETANI_MEMBUDIDAYAKAN_IKAN definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PETANI_MEMBUDIDAYAKAN_IKAN;

CREATE TABLE PETANI_MEMBUDIDAYAKAN_IKAN (
	KdPetani int NOT NULL,
	NamaTambak varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	Lokasi varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	StokDewasa int NULL,
	BiayaPerSatuan int NULL,
	Keuntungan int NULL,
	KdIkan int NOT NULL,
	CONSTRAINT PK_PETANI_MEMBUDIDAYAKAN_IKAN PRIMARY KEY (KdPetani,KdIkan),
	CONSTRAINT FK__PETANI_MEMBUDIDAYAKAN_IKAN__KdIkan FOREIGN KEY (KdIkan) REFERENCES IKAN(KdIkan),
	CONSTRAINT FK__PETANI_MEMBUDIDAYAKAN_IKAN__KdPetani FOREIGN KEY (KdPetani) REFERENCES PETANI(KdPetani)
);


-- DBMSProject.dbo.PRODUK definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PRODUK;

CREATE TABLE PRODUK (
	KdProduk int NOT NULL,
	Nama varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	Deskripsi varchar(1000) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	KdIkan int NULL,
	CONSTRAINT PK_PRODUK PRIMARY KEY (KdProduk),
	CONSTRAINT FK__PRODUK__KdIkan FOREIGN KEY (KdIkan) REFERENCES IKAN(KdIkan)
);


-- DBMSProject.dbo.SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK;

CREATE TABLE SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK (
	KdSupplier int NOT NULL,
	KdProduk int NOT NULL,
	KdPabrik int NOT NULL,
	Tanggal date NOT NULL,
	Jumlah int NULL,
	CONSTRAINT PK_SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK PRIMARY KEY (KdSupplier,KdProduk,KdPabrik,Tanggal),
	CONSTRAINT FK__SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK__KdPabrik FOREIGN KEY (KdPabrik) REFERENCES PABRIK(KdPabrik),
	CONSTRAINT FK__SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK__KdProduk FOREIGN KEY (KdProduk) REFERENCES PRODUK(KdProduk),
	CONSTRAINT FK__SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK__KdSupplier FOREIGN KEY (KdSupplier) REFERENCES SUPPLIER(KdSupplier)
);


-- DBMSProject.dbo.SUPPLIER_MENSUPLAI_PRODUK definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.SUPPLIER_MENSUPLAI_PRODUK;

CREATE TABLE SUPPLIER_MENSUPLAI_PRODUK (
	KdSupplier int NOT NULL,
	KdProduk int NOT NULL,
	BiayaOperasional int NULL,
	Keuntungan int NULL,
	Stok int NULL,
	CONSTRAINT PK_SUPPLIER_MENSUPLAI_PRODUK PRIMARY KEY (KdSupplier,KdProduk),
	CONSTRAINT FK__SUPPLIER_MENSUPLAI_PRODUK__KdProduk FOREIGN KEY (KdProduk) REFERENCES PRODUK(KdProduk),
	CONSTRAINT FK__SUPPLIER_MENSUPLAI_PRODUK__KdSupplier FOREIGN KEY (KdSupplier) REFERENCES SUPPLIER(KdSupplier)
);


-- DBMSProject.dbo.PABRIK_MENGHASILKAN_PRODUK definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PABRIK_MENGHASILKAN_PRODUK;

CREATE TABLE PABRIK_MENGHASILKAN_PRODUK (
	KdPabrik int NOT NULL,
	KdProduk int NOT NULL,
	BiayaProduksi int NULL,
	Keuntungan int NULL,
	Stok int NULL,
	CONSTRAINT PK_PABRIK_MENGHASILKAN_PRODUK PRIMARY KEY (KdPabrik,KdProduk),
	CONSTRAINT FK__PABRIK_MENGHASILKAN_PRODUK__KdPabrik FOREIGN KEY (KdPabrik) REFERENCES PABRIK(KdPabrik),
	CONSTRAINT FK__PABRIK_MENGHASILKAN_PRODUK__KdProduk FOREIGN KEY (KdProduk) REFERENCES PRODUK(KdProduk)
);


-- DBMSProject.dbo.PELANGGAN_MEMBELI_PRODUK_DARI_SUPPLIER definition

-- Drop table

-- DROP TABLE DBMSProject.dbo.PELANGGAN_MEMBELI_PRODUK_DARI_SUPPLIER;

CREATE TABLE PELANGGAN_MEMBELI_PRODUK_DARI_SUPPLIER (
	KdSupplier int NOT NULL,
	KdProduk int NOT NULL,
	Jumlah int NULL,
	KdPelanggan int NOT NULL,
	Tanggal date NOT NULL,
	CONSTRAINT PK_PELANGGAN_MEMBELI_PRODUK_DARI_SUPPLIER PRIMARY KEY (KdPelanggan,KdProduk,KdSupplier,Tanggal),
	CONSTRAINT FK__PELANGGAN_MEMBELI_PRODUK__KdPelanggan FOREIGN KEY (KdPelanggan) REFERENCES PELANGGAN(KdPelanggan),
	CONSTRAINT FK__PELANGGAN_MEMBELI_PRODUK__KdProduk FOREIGN KEY (KdProduk) REFERENCES PRODUK(KdProduk),
	CONSTRAINT FK__PELANGGAN_MEMBELI_PRODUK__KdSupplier FOREIGN KEY (KdSupplier) REFERENCES SUPPLIER(KdSupplier)
);


-- dbo.V_BUKU_BelanjaPabrik source

-- Tampilan total belanja ikan dari nelayan dan petani tambak oleh pabrik
CREATE VIEW V_BUKU_BelanjaPabrik AS
SELECT KdPabrik, Nama, SUM(BelanjaTotal) BelanjaTotal FROM (
	SELECT P.KdPabrik, P.Nama, ISNULL(SUM(PMIDN.Jumlah * (I.HargaJual + NMI.Keuntungan)), 0) BelanjaTotal
	FROM PABRIK P
	LEFT JOIN PABRIK_MEMBELI_IKAN_DARI_NELAYAN PMIDN ON PMIDN.KdPabrik = P.KdPabrik
	LEFT JOIN IKAN I ON I.KdIkan = PMIDN.KdIkan
	LEFT JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdNelayan = PMIDN.KdNelayan AND NMI.KdIkan = PMIDN.KdIkan
	GROUP BY P.KdPabrik, P.Nama
	UNION
	SELECT P.KdPabrik, P.Nama, ISNULL(SUM(PMIDP.Jumlah * (I.HargaJual + PMI.BiayaPerSatuan + PMI.Keuntungan)), 0) BelanjaTotal
	FROM PABRIK P
	LEFT JOIN PABRIK_MEMBELI_IKAN_DARI_PETANI PMIDP ON PMIDP.KdPabrik = P.KdPabrik
	LEFT JOIN IKAN I ON I.KdIkan = PMIDP.KdIkan
	LEFT JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdPetani = PMIDP.KdPetani AND PMI.KdIkan = PMIDP.KdIkan
	GROUP BY P.KdPabrik, P.Nama
) BELANJA GROUP BY KdPabrik, Nama;


-- dbo.V_BUKU_BelanjaPelanggan source

-- Tampilan total belanja produk dari supplier oleh pelanggan
CREATE VIEW V_BUKU_BelanjaPelanggan AS
SELECT P.KdPelanggan, P.Nama, ISNULL(SUM(PMPDS.Jumlah * (I.HargaJual + SMP.BiayaOperasional + SMP.Keuntungan)), 0) BelanjaTotal
FROM PELANGGAN P
LEFT JOIN PELANGGAN_MEMBELI_PRODUK_DARI_SUPPLIER PMPDS ON PMPDS.KdPelanggan = P.KdPelanggan
LEFT JOIN PRODUK PR ON PR.KdProduk = PMPDS.KdProduk
LEFT JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdSupplier = PMPDS.KdSupplier AND SMP.KdProduk = PMPDS.KdProduk
LEFT JOIN IKAN I ON I.KdIkan = PR.KdIkan
GROUP BY P.KdPelanggan, P.Nama;


-- dbo.V_BUKU_BelanjaPetani source

-- Tampilan total belanja ikan dari nelayan oleh petani tambak
CREATE VIEW V_BUKU_BelanjaPetani AS
SELECT P.KdPetani, P.Nama, ISNULL(SUM(PMIDN.Jumlah * (I.HargaJual + NMI.Keuntungan)), 0) BelanjaTotal
FROM PETANI P
LEFT JOIN PETANI_MEMBELI_IKAN_DARI_NELAYAN PMIDN ON PMIDN.KdPetani = P.KdPetani
LEFT JOIN IKAN I ON I.KdIkan = PMIDN.KdIkan
LEFT JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdNelayan = PMIDN.KdNelayan AND NMI.KdIkan = PMIDN.KdIkan
GROUP BY P.KdPetani, P.Nama;


-- dbo.V_BUKU_BelanjaSupplier source

-- Tampilan total belanja produk dari pabrik oleh supplier
CREATE VIEW V_BUKU_BelanjaSupplier AS
SELECT S.KdSupplier, S.Nama, ISNULL(SUM(SMPDP.Jumlah * (I.HargaJual + PMP.BiayaProduksi + PMP.Keuntungan)), 0) BelanjaTotal
FROM SUPPLIER S
LEFT JOIN SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK SMPDP ON SMPDP.KdSupplier = S.KdSupplier
LEFT JOIN PRODUK P ON P.KdProduk = SMPDP.KdProduk
LEFT JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdPabrik = SMPDP.KdPabrik AND PMP.KdProduk = SMPDP.KdProduk
LEFT JOIN IKAN I ON I.KdIkan = P.KdIkan
GROUP BY S.KdSupplier, S.Nama;


-- dbo.V_BUKU_PendapatanNelayan source

-- Tampilan total pendapatan nelayan dalam penjualan ikan ke petani tambak dan pabrik
CREATE VIEW V_BUKU_PendapatanNelayan AS
SELECT KdNelayan, Nama, SUM(PendapatanTotal) PendapatanTotal FROM (
	SELECT N.KdNelayan, N.Nama, ISNULL(SUM(PMIDN.Jumlah * (I.HargaJual + NMI.Keuntungan)), 0) PendapatanTotal
	FROM NELAYAN N
	LEFT JOIN PABRIK_MEMBELI_IKAN_DARI_NELAYAN PMIDN ON PMIDN.KdNelayan = N.KdNelayan
	LEFT JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdNelayan = PMIDN.KdNelayan AND NMI.KdIkan = PMIDN.KdIkan
	LEFT JOIN IKAN I ON I.KdIkan = PMIDN.KdIkan
	GROUP BY N.KdNelayan, N.Nama
	UNION
	SELECT N.KdNelayan, N.Nama, ISNULL(SUM(PMIDN.Jumlah * (I.HargaJual + NMI.Keuntungan)), 0) PendapatanTotal
	FROM NELAYAN N
	LEFT JOIN PETANI_MEMBELI_IKAN_DARI_NELAYAN PMIDN ON PMIDN.KdNelayan = N.KdNelayan
	LEFT JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdNelayan = PMIDN.KdNelayan AND NMI.KdIkan = PMIDN.KdIkan
	LEFT JOIN IKAN I ON I.KdIkan = PMIDN.KdIkan
	GROUP BY N.KdNelayan, N.Nama
) PENDAPATAN GROUP BY KdNelayan, Nama;


-- dbo.V_BUKU_PendapatanPabrik source

-- Tampilan total pendapatan pabrik dalam penjualan produk ke supplier
CREATE VIEW V_BUKU_PendapatanPabrik AS
SELECT P.KdPabrik, P.Nama, ISNULL(SUM(SMPDP.Jumlah * (I.HargaJual + PMP.BiayaProduksi + PMP.Keuntungan)), 0) PendapatanTotal
FROM PABRIK P
LEFT JOIN SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK SMPDP ON SMPDP.KdPabrik = P.KdPabrik
LEFT JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdPabrik = SMPDP.KdPabrik AND PMP.KdProduk = SMPDP.KdProduk
LEFT JOIN PRODUK PR ON PR.KdProduk = SMPDP.KdProduk
LEFT JOIN IKAN I ON I.KdIkan = PR.KdIkan
GROUP BY P.KdPabrik, P.Nama;


-- dbo.V_BUKU_PendapatanPetani source

-- Tampilan total pendapatan petani tambak dalam penjualan ikan ke pabrik
CREATE VIEW V_BUKU_PendapatanPetani AS
SELECT P.KdPetani, P.Nama, ISNULL(SUM(PMIDP.Jumlah * (I.HargaJual + PMI.BiayaPerSatuan + PMI.Keuntungan)), 0) PendapatanTotal
FROM PETANI P
LEFT JOIN PABRIK_MEMBELI_IKAN_DARI_PETANI PMIDP ON PMIDP.KdPetani = P.KdPetani
LEFT JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdPetani = PMIDP.KdPetani AND PMI.KdIkan = PMIDP.KdIkan
LEFT JOIN IKAN I ON I.KdIkan = PMIDP.KdIkan
GROUP BY P.KdPetani, P.Nama;


-- dbo.V_BUKU_PendapatanSupplier source

-- Tampilan total pendapatan supplier dalam penjualan produk ke pelanggan
CREATE VIEW V_BUKU_PendapatanSupplier AS
SELECT S.KdSupplier, S.Nama, ISNULL(SUM(PMPDS.Jumlah * (I.HargaJual + SMP.BiayaOperasional + SMP.Keuntungan)), 0) PendapatanTotal
FROM SUPPLIER S
LEFT JOIN PELANGGAN_MEMBELI_PRODUK_DARI_SUPPLIER PMPDS ON PMPDS.KdSupplier = S.KdSupplier
LEFT JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdSupplier = PMPDS.KdSupplier AND SMP.KdProduk = PMPDS.KdProduk
LEFT JOIN PRODUK P ON P.KdProduk = PMPDS.KdProduk
LEFT JOIN IKAN I ON I.KdIkan = P.KdIkan
GROUP BY S.KdSupplier, S.Nama;


-- dbo.V_IKAN_DiLaut source

-- Tampilan sebaran spesies ikan di laut
CREATE VIEW V_IKAN_DiLaut AS
SELECT I.Nama, L.KdLaut, L.Nama Laut
FROM IKAN I
JOIN LAUT_MEMILIKI_IKAN LMI ON LMI.KdIkan = I.KdIkan
JOIN LAUT L ON L.KdLaut = LMI.KdLaut;


-- dbo.V_IKAN_Dibudidayakan source

-- Tampilan daftar ikan yang dibudidayakan
CREATE VIEW V_IKAN_Dibudidayakan AS
SELECT I.KdIkan, I.Nama, COUNT(*) AS JmlTambak
FROM IKAN I
JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdIkan = I.KdIkan
GROUP BY I.KdIkan, I.Nama;


-- dbo.V_IKAN_JumlahPerLaut source

-- Tampilan sebaran jumlah spesies ikan per laut
CREATE VIEW V_IKAN_JumlahPerLaut AS
SELECT L.KdLaut, L.Nama, COUNT(*) JmlSpesies
FROM LAUT L
JOIN LAUT_MEMILIKI_IKAN LMI ON LMI.KdLaut = L.KdLaut
GROUP BY L.KdLaut, L.Nama;


-- dbo.V_JELAJAH_IkanTambak source

-- Tampilan daftar ikan beserta harga jualnya yang dibudidayakan petani tambak
CREATE VIEW V_JELAJAH_IkanTambak AS
SELECT P.KdPetani, P.Nama Petani, I.KdIkan, I.Nama, (I.HargaJual + PMI.BiayaPerSatuan + PMI.Keuntungan) Harga, PMI.StokDewasa
FROM IKAN I
JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdIkan = I.KdIkan
JOIN PETANI P ON P.KdPetani = PMI.KdPetani;


-- dbo.V_JELAJAH_IkanTambakTermurah source

-- Tampilan harga (termurah sampai termahal) dan stok dewasa ikan dari tambak yang dapat dibeli pabrik
CREATE VIEW V_JELAJAH_IkanTambakTermurah AS
SELECT I.KdIkan, I.Nama, P.KdPetani, P.Nama NamaPetani, (I.HargaJual + PMI.BiayaPerSatuan + PMI.Keuntungan) Harga, PMI.StokDewasa
FROM IKAN I
JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdIkan = I.KdIkan
JOIN PETANI P ON P.KdPetani = PMI.KdPetani;


-- dbo.V_JELAJAH_IkanTangkapan source

-- Tampilan daftar ikan beserta harga jualnya yang ditangkap nelayan
CREATE VIEW V_JELAJAH_IkanTangkapan AS
SELECT N.KdNelayan, N.Nama Nelayan, I.KdIkan, I.Nama, (I.HargaJual + NMI.Keuntungan) Harga, NMI.Persediaan
FROM IKAN I
JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdIkan = I.KdIkan
JOIN NELAYAN N ON N.KdNelayan = NMI.KdNelayan;


-- dbo.V_JELAJAH_IkanTangkapanTermurah source

-- Tampilan harga (termurah sampai termahal) dan persediaan ikan dari nelayan yang dapat dibeli pabrik dan petani tambak
CREATE VIEW V_JELAJAH_IkanTangkapanTermurah AS
SELECT I.KdIkan, I.Nama, N.KdNelayan, N.Nama NamaNelayan, (I.HargaJual + NMI.Keuntungan) Harga, NMI.Persediaan
FROM IKAN I
JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdIkan = I.KdIkan
JOIN NELAYAN N ON N.KdNelayan = NMI.KdNelayan;


-- dbo.V_JELAJAH_ProdukPabrik source

-- Tampilan daftar produk beserta harga jualnya yang dihasilkan pabrik
CREATE VIEW V_JELAJAH_ProdukPabrik AS
SELECT PA.KdPabrik, PA.Nama Pabrik, P.KdProduk, P.Nama, (I.HargaJual + PMP.BiayaProduksi + PMP.Keuntungan) Harga, PMP.Stok
FROM PRODUK P
JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdProduk = P.KdProduk
JOIN PABRIK PA ON PA.KdPabrik = PMP.KdPabrik
JOIN IKAN I ON I.KdIkan = P.KdIkan;


-- dbo.V_JELAJAH_ProdukPabrikTermurah source

-- Tampilan harga (termurah sampai termahal) dan stok produk dari pabrik yang dapat dibeli supplier
CREATE VIEW V_JELAJAH_ProdukPabrikTermurah AS
SELECT P.KdProduk, P.Nama, PA.KdPabrik, PA.Nama NamaPabrik, (I.HargaJual + PMP.BiayaProduksi + PMP.Keuntungan) Harga, PMP.Stok
FROM PRODUK P
JOIN IKAN I ON I.KdIkan = P.KdIkan
JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdProduk = P.KdProduk
JOIN PABRIK PA ON PA.KdPabrik = PMP.KdPabrik;


-- dbo.V_JELAJAH_ProdukSupplier source

-- Tampilan daftar produk beserta harga jualnya yang disediakan supplier
CREATE VIEW V_JELAJAH_ProdukSupplier AS
SELECT S.KdSupplier, S.Nama Supplier, P.KdProduk, P.Nama, (I.HargaJual + SMP.BiayaOperasional + SMP.Keuntungan) Harga, SMP.Stok
FROM PRODUK P
JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdProduk = P.KdProduk
JOIN SUPPLIER S ON S.KdSupplier = SMP.KdSupplier
JOIN IKAN I ON I.KdIkan = P.KdIkan;


-- dbo.V_JELAJAH_ProdukSupplierTermurah source

-- Tampilan harga (termurah sampai termahal) dan stok produk dari supplier yang dapat dibeli pelanggan
CREATE VIEW V_JELAJAH_ProdukSupplierTermurah AS
SELECT P.KdProduk, P.Nama, S.KdSupplier, S.Nama NamaSupplier, (I.HargaJual + SMP.BiayaOperasional + SMP.Keuntungan) Harga, SMP.Stok
FROM PRODUK P
JOIN IKAN I ON I.KdIkan = P.KdIkan
JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdProduk = P.KdProduk
JOIN SUPPLIER S ON S.KdSupplier = SMP.KdSupplier;


-- dbo.V_KAPAL_DigunakanTanggal source

-- Tampilan daftar nelayan yang menggunakan kapal tertentu pada tanggal tertentu
CREATE VIEW V_KAPAL_DigunakanTanggal AS
SELECT NMK.KdKapal, K.Nama Kapal, NMK.Tanggal, N.KdNelayan, N.Nama Nelayan
FROM NELAYAN_MENGGUNAKAN_KAPAL NMK
JOIN NELAYAN N ON N.KdNelayan = NMK.KdNelayan
JOIN KAPAL K ON K.KdKapal = NMK.KdKapal;


-- dbo.V_KAPAL_Overkapasitas source

-- Tampilan pada tanggal berapa suatu kapal overkapasitas
CREATE VIEW V_KAPAL_Overkapasitas AS
SELECT NMK.Tanggal, K.KdKapal, K.Nama, K.Kapasitas, COUNT(*) JmlAwak
FROM KAPAL K
JOIN NELAYAN_MENGGUNAKAN_KAPAL NMK ON NMK.KdKapal = K.KdKapal
GROUP BY K.KdKapal, K.Nama, K.Kapasitas, NMK.Tanggal HAVING COUNT(*) > K.Kapasitas;


-- dbo.V_KAPAL_Penggunaan source

-- Tampilan jumlah penggunaan kapal
CREATE VIEW V_KAPAL_Penggunaan AS
SELECT KdKapal, Nama, COUNT(*) JmlHari FROM (
	SELECT K.KdKapal, K.Nama, NMK.Tanggal, COUNT(*) AS JmlTerpakai
	FROM KAPAL K
	JOIN NELAYAN_MENGGUNAKAN_KAPAL NMK ON NMK.KdKapal = K.KdKapal
	GROUP BY K.KdKapal, K.Nama, NMK.Tanggal
) TGLPENGGUNAAN GROUP BY KdKapal, Nama;


-- dbo.V_KAPAL_PenggunaanTanggal source

-- Tampilan daftar nelayan yang berlayar pada tanggal tertentu
CREATE VIEW V_KAPAL_PenggunaanTanggal AS
SELECT NMK.Tanggal, K.KdKapal, K.Nama Kapal, N.KdNelayan, N.Nama Nelayan
FROM NELAYAN_MENGGUNAKAN_KAPAL NMK
JOIN NELAYAN N ON N.KdNelayan = NMK.KdNelayan
JOIN KAPAL K ON K.KdKapal = NMK.KdKapal;


-- dbo.V_LIST_Ikan source

-- Tampilan daftar ikan
CREATE VIEW V_LIST_Ikan AS
SELECT KdIkan, Nama, HargaJual FROM IKAN;


-- dbo.V_LIST_Kapal source

-- Tampilan daftar kapal beserta lokasinya di laut
CREATE VIEW V_LIST_Kapal AS
SELECT K.KdKapal, K.Nama, K.Jenis, K.Kapasitas, L.Nama Lokasi
FROM KAPAL K
JOIN LAUT L ON L.KdLaut = K.KdLaut;


-- dbo.V_LIST_Laut source

-- Tampilan daftar laut
CREATE VIEW V_LIST_Laut AS
SELECT KdLaut, Nama, Lokasi FROM LAUT;


-- dbo.V_LIST_Nelayan source

-- Tampilan daftar nelayan
CREATE VIEW V_LIST_Nelayan AS
SELECT KdNelayan, Nama, NoTelp FROM NELAYAN;


-- dbo.V_LIST_Pabrik source

-- Tampilan daftar pabrik
CREATE VIEW V_LIST_Pabrik AS
SELECT KdPabrik, Nama, Lokasi, NoTelp FROM PABRIK;


-- dbo.V_LIST_Pelanggan source

-- Tampilan daftar pelanggan
CREATE VIEW V_LIST_Pelanggan AS
SELECT KdPelanggan, Nama, Lokasi FROM PELANGGAN;


-- dbo.V_LIST_Petani source

-- Tampilan daftar petani
CREATE VIEW V_LIST_Petani AS
SELECT KdPetani, Nama, NoTelp FROM PETANI;


-- dbo.V_LIST_Produk source

-- Tampilan daftar produk beserta ikan asalnya
CREATE VIEW V_LIST_Produk AS
SELECT P.KdProduk, P.Nama, P.Deskripsi, I.Nama AsalIkan
FROM PRODUK P
JOIN IKAN I ON I.KdIkan = P.KdIkan;


-- dbo.V_LIST_Supplier source

-- Tampilan daftar supplier
CREATE VIEW V_LIST_Supplier AS
SELECT KdSupplier, Nama, Lokasi, NoTelp FROM SUPPLIER;


-- dbo.V_STATS_RankNelayan source

-- Tampilan ranking nelayan berdasarkan nilai seluruh persediaan tangkapan ikan
CREATE VIEW V_STATS_RankNelayan AS
SELECT N.KdNelayan, N.Nama, ISNULL(SUM(I.HargaJual * NMI.Persediaan), 0) Nilai
FROM NELAYAN N
LEFT JOIN NELAYAN_MENANGKAP_IKAN NMI ON NMI.KdNelayan = N.KdNelayan
LEFT JOIN IKAN I ON I.KdIkan = NMI.KdIkan
GROUP BY N.KdNelayan, N.Nama;


-- dbo.V_STATS_RankPabrik source

-- Tampilan ranking pabrik berdasarkan nilai seluruh stok produk yang dihasilkan
CREATE VIEW V_STATS_RankPabrik AS
SELECT P.KdPabrik, P.Nama, ISNULL(SUM(I.HargaJual * PMP.Stok), 0) Nilai
FROM PABRIK P
LEFT JOIN PABRIK_MENGHASILKAN_PRODUK PMP ON PMP.KdPabrik = P.KdPabrik
LEFT JOIN PRODUK PR ON PR.KdProduk = PMP.KdProduk
LEFT JOIN IKAN I ON I.KdIkan = PR.KdIkan
GROUP BY P.KdPabrik, P.Nama;


-- dbo.V_STATS_RankPetani source

-- Tampilan ranking petani tambak berdasarkan nilai seluruh stok dewasa ikan
CREATE VIEW V_STATS_RankPetani AS
SELECT P.KdPetani, P.Nama, ISNULL(SUM(I.HargaJual * PMI.StokDewasa), 0) Nilai
FROM PETANI P
LEFT JOIN PETANI_MEMBUDIDAYAKAN_IKAN PMI ON PMI.KdPetani = P.KdPetani
LEFT JOIN IKAN I ON I.KdIkan = PMI.KdIkan
GROUP BY P.KdPetani, P.Nama;


-- dbo.V_STATS_RankSupplier source

-- Tampilan ranking supplier berdasarkan nilai seluruh stok produk yang disediakan
CREATE VIEW V_STATS_RankSupplier AS
SELECT S.KdSupplier, S.Nama, ISNULL(SUM(I.HargaJual * SMP.Stok), 0) Nilai
FROM SUPPLIER S
LEFT JOIN SUPPLIER_MENSUPLAI_PRODUK SMP ON SMP.KdSupplier = S.KdSupplier
LEFT JOIN PRODUK PR ON PR.KdProduk = SMP.KdProduk
LEFT JOIN IKAN I ON I.KdIkan = PR.KdIkan
GROUP BY S.KdSupplier, S.Nama;


-- dbo.V_TERDEKAT_Pabrik source

-- Tampilan pabrik terdekat dengan supplier
CREATE VIEW V_TERDEKAT_Pabrik AS
SELECT S.KdSupplier, S.Nama NamaSupplier, P.Lokasi, P.KdPabrik, P.Nama NamaPabrik
FROM PABRIK P
JOIN SUPPLIER S ON S.Lokasi = P.Lokasi;


-- dbo.V_TERDEKAT_Supplier source

-- Tampilan supplier terdekat dengan pelanggan
CREATE VIEW V_TERDEKAT_Supplier AS
SELECT P.KdPelanggan, P.Nama NamaPelanggan, S.Lokasi, S.KdSupplier, S.Nama NamaSupplier
FROM SUPPLIER S
JOIN PELANGGAN P ON P.Lokasi = S.Lokasi;


-- dbo.V_TERDEKAT_Tambak source

-- Tampilan tambak terdekat dengan pabrik
CREATE VIEW V_TERDEKAT_Tambak AS
SELECT P.KdPabrik, P.Nama NamaPabrik, PMI.Lokasi, PMI.KdPetani, PMI.KdIkan, I.Nama NamaIkan, PMI.NamaTambak NamaTambak
FROM PETANI_MEMBUDIDAYAKAN_IKAN PMI
JOIN PABRIK P ON P.Lokasi = PMI.Lokasi
JOIN IKAN I ON I.KdIkan = PMI.KdIkan;

EXEC sp_MSforeachtable "ALTER TABLE ? NOCHECK CONSTRAINT ALL"; -- Disable FK checks

INSERT INTO IKAN (KdIkan,Nama,HargaJual) VALUES
(1,N'Tongkol',3000),
(2,N'Tuna',15000),
(3,N'Kakap',10000),
(4,N'Makarel',2000),
(5,N'Tenggiri',5000),
(6,N'Kerang',4000),
(7,N'Kerang Mutiara',50000),
(8,N'Cumi-Cumi',6000),
(9,N'Kepiting',8000),
(10,N'Udang',100);
INSERT INTO KAPAL (KdKapal,Nama,Jenis,Kapasitas,KdLaut) VALUES
(1,N'Mercury',N'Perahu',3,2),
(2,N'Argon',N'Trawler',5,2),
(3,N'Neon',N'Seiner',4,3),
(4,N'Oxygen',N'Longliner',7,1),
(5,N'Radium',N'Trawler',2,3);
INSERT INTO LAUT (KdLaut,Nama,Lokasi) VALUES
(1,N'Laut Surabaya',N'Surabaya'),
(2,N'Laut Banyuwangi',N'Banyuwangi'),
(3,N'Laut Pasuruan',N'Pasuruan');
INSERT INTO LAUT_MEMILIKI_IKAN (KdLaut,KdIkan) VALUES
(1,1),
(1,2),
(1,4),
(1,7),
(1,8),
(1,9),
(1,10),
(2,2),
(2,3),
(2,5),
(2,6),
(2,8),
(2,10),
(3,1),
(3,2),
(3,3),
(3,4),
(3,6),
(3,9),
(3,10);
INSERT INTO NELAYAN (KdNelayan,Nama,NoTelp) VALUES
(1,N'Kasno',N'088869329842'),
(2,N'Suparjo',N'087778529232'),
(3,N'Sunarko',N'088359424803'),
(4,N'Winarko',N'085707872273'),
(5,N'Agung Hadibroto',N'080902663037'),
(6,N'Edi Kusumorahardjo',N'084828957377'),
(7,N'Agus Setiabudi',N'086314174935'),
(8,N'Bagus Sunarto',N'086418340174'),
(9,N'Wisnu Atmaja',N'085196272811'),
(10,N'Putra Sutanto',N'087943669497');
INSERT INTO NELAYAN_MENANGKAP_IKAN (KdNelayan,KdIkan,Keuntungan,Persediaan) VALUES
(1,1,500,52),
(1,3,1000,24),
(1,5,500,10),
(2,2,2000,29),
(2,4,500,56),
(2,6,1000,32),
(3,7,10000,12),
(3,8,1000,34),
(4,9,2000,26),
(4,10,400,54),
(5,1,2000,32),
(5,2,5000,12),
(5,3,5000,14),
(6,4,3000,78),
(6,5,1000,24),
(6,6,1000,57),
(7,7,5000,43),
(7,9,1000,30),
(8,8,2000,66),
(8,10,100,87);
INSERT INTO NELAYAN_MENGGUNAKAN_KAPAL (KdNelayan,KdKapal,Tanggal) VALUES
(1,1,'2021-10-03'),
(1,1,'2021-10-10'),
(1,3,'2021-10-06'),
(2,1,'2021-10-03'),
(2,1,'2021-10-06'),
(3,5,'2021-10-04'),
(4,3,'2021-10-07'),
(4,5,'2021-10-04'),
(5,5,'2021-10-04'),
(6,2,'2021-10-05'),
(6,4,'2021-10-09'),
(7,3,'2021-10-05'),
(7,5,'2021-10-08'),
(8,4,'2021-10-05'),
(9,4,'2021-10-05'),
(10,1,'2021-10-08');
INSERT INTO PABRIK (KdPabrik,Nama,Lokasi,NoTelp) VALUES
(1,N'PT. Alam Jaya',N'Surabaya',N'089260075433'),
(2,N'PT. Amarta Mandiri',N'Surabaya',N'082137399274'),
(3,N'PT. Asia Sejahtera Mina',N'Sidoarjo',N'084956530030'),
(4,N'UD. Ananta Pratama',N'Banyuwangi',N'080179017885'),
(5,N'PT. Agromina Wicaksana',N'Sidoarjo',N'083019667020'),
(6,N'PT. Alter Trade Indonesia',N'Sidoarjo',N'082161568067'),
(7,N'PT. Aneka Tuna Indonesia',N'Pasuruan',N'085652173427'),
(8,N'PT. Aneka Sumber Alam Jaya',N'Pasuruan',N'089428524284'),
(9,N'PT. Agro Mitra Alimentare',N'Malang',N'083846371580'),
(10,N'PT. Avila Prima Intra Makmur',N'Banyuwangi',N'080922562008'),
(11,N'CV. Amin Jaya Makmur',N'Sidoarjo',N'085210751770');
INSERT INTO PABRIK_MEMBELI_IKAN_DARI_NELAYAN (KdPabrik,KdIkan,KdNelayan,Jumlah,Tanggal) VALUES
(1,1,5,10,'2021-09-01'),
(1,2,5,5,'2021-09-01'),
(1,4,2,20,'2021-09-02'),
(2,9,7,4,'2021-09-05'),
(3,3,1,5,'2021-09-05'),
(4,5,1,12,'2021-09-08'),
(4,6,6,21,'2021-09-10');
INSERT INTO PABRIK_MEMBELI_IKAN_DARI_PETANI (KdPabrik,KdIkan,KdPetani,Jumlah,Tanggal) VALUES
(8,4,1,500,'2021-09-02'),
(8,9,2,300,'2021-09-04');
INSERT INTO PABRIK_MENGHASILKAN_PRODUK (KdPabrik,KdProduk,BiayaProduksi,Keuntungan,Stok) VALUES
(1,1,1000,500,500),
(1,2,1000,2000,350),
(1,4,1000,500,700),
(1,7,1000,1000,600),
(1,8,1000,1000,400),
(1,9,1000,100,1000),
(2,12,1000,2000,700),
(2,13,10000,20000,150),
(2,14,500,2000,800),
(3,1,1000,500,750),
(3,3,1000,1000,700),
(3,4,1000,500,800),
(4,2,2000,3000,300),
(4,3,2000,3000,350),
(4,5,2000,1000,650),
(4,6,2000,1000,750),
(5,13,10000,10000,300),
(6,2,2000,3000,1400),
(7,2,3000,2000,800),
(7,3,3000,2000,1200),
(7,10,3000,2000,1000),
(8,1,1000,1000,500),
(8,4,1000,1000,600),
(8,6,1000,1000,650),
(8,8,1000,1000,550),
(8,9,100,500,1150),
(8,12,1000,1000,450),
(8,14,1000,1000,400),
(9,5,1500,1500,600),
(9,6,1500,1500,600),
(9,7,1500,1500,600),
(9,8,1500,1500,600),
(10,7,500,2500,450),
(10,9,50,50,1500),
(10,10,1000,2000,400),
(10,11,500,1500,550),
(11,14,1000,1000,1250);
INSERT INTO PELANGGAN (KdPelanggan,Nama,Lokasi) VALUES
(1,N'Superindo Surabaya',N'Surabaya'),
(2,N'Indomaret Pasuruan',N'Pasuruan'),
(3,N'RM Seafood Pak Jono Banyuwangi',N'Banyuwangi'),
(4,N'Hotel Seashore Pasuruan',N'Pasuruan'),
(5,N'Pasar Wage Malang',N'Malang'),
(6,N'RM Agro Wicaksana Sidoarjo',N'Sidoarjo'),
(7,N'Alfamidi Surabaya',N'Surabaya'),
(8,N'Hypermart Banyuwangi',N'Banyuwangi');
INSERT INTO PELANGGAN_MEMBELI_PRODUK_DARI_SUPPLIER (KdSupplier,KdProduk,Jumlah,KdPelanggan,Tanggal) VALUES
(1,5,100,1,'2021-11-01'),
(1,6,100,1,'2021-11-02'),
(2,8,100,1,'2021-11-03'),
(3,1,300,2,'2021-11-02'),
(3,2,100,2,'2021-11-03'),
(4,10,120,2,'2021-11-04'),
(4,14,250,2,'2021-11-05'),
(5,2,100,3,'2021-11-03'),
(5,7,350,3,'2021-11-04'),
(6,3,100,4,'2021-11-04'),
(6,9,250,4,'2021-11-05'),
(6,11,300,4,'2021-11-06'),
(8,1,300,5,'2021-11-07'),
(7,2,50,5,'2021-11-05'),
(8,12,200,5,'2021-11-08'),
(7,14,300,5,'2021-11-06'),
(9,1,350,6,'2021-11-06'),
(10,2,100,6,'2021-11-07'),
(10,4,400,6,'2021-11-08'),
(1,6,150,7,'2021-11-07'),
(2,7,150,7,'2021-11-08'),
(2,8,100,7,'2021-11-09'),
(3,1,250,8,'2021-11-08'),
(3,2,150,8,'2021-11-09'),
(3,3,200,8,'2021-11-10'),
(4,12,50,8,'2021-11-11');
INSERT INTO PETANI (KdPetani,Nama,NoTelp) VALUES
(1,N'Erik Hanjoyo',N'081010642385'),
(2,N'Wahyu Gunawan',N'081222600775'),
(3,N'Bambang Wibisono',N'089394019751');
INSERT INTO PETANI_MEMBELI_IKAN_DARI_NELAYAN (KdPetani,KdIkan,KdNelayan,Jumlah,Tanggal) VALUES
(1,1,1,25,'2021-09-15'),
(1,10,4,10,'2021-09-16'),
(3,8,3,19,'2021-09-20');
INSERT INTO PETANI_MEMBUDIDAYAKAN_IKAN (KdPetani,NamaTambak,Lokasi,StokDewasa,BiayaPerSatuan,Keuntungan,KdIkan) VALUES
(1,N'Tongkol Pak Hanjoyo',N'Pasuruan',2000,500,500,1),
(1,N'Makarel Pak Hanjoyo',N'Pasuruan',3500,750,750,4),
(1,N'Udang Pak Hanjoyo',N'Banyuwangi',7000,50,50,10),
(2,N'Crabfarm Mr. Gunawan',N'Surabaya',10000,500,500,9),
(3,N'Cumi Wibi Aquaculture',N'Banyuwangi',5000,1000,500,8),
(3,N'Udang Wibi Aquaculture',N'Surabaya',8000,100,100,10);
INSERT INTO PRODUK (KdProduk,Nama,Deskripsi,KdIkan) VALUES
(1,N'Tongkol Beku',N'Daging beku ikan tongkol',1),
(2,N'Tuna Beku',N'Daging beku ikan tuna',2),
(3,N'Kakap Beku',N'Daging beku ikan kakap',3),
(4,N'Makarel Beku',N'Daging beku ikan makarel',4),
(5,N'Tenggiri Beku',N'Daging beku ikan tenggiri',5),
(6,N'Kerang Beku',N'Kulit dan daging beku kerang',6),
(7,N'Cumi-Cumi Beku',N'Daging beku cumi-cumi',8),
(8,N'Kepiting Beku',N'Daging beku kepiting',9),
(9,N'Udang Beku',N'Daging beku udang',10),
(10,N'Kakap Segar',N'Daging segar ikan kakap',3),
(11,N'Tenggiri Segar',N'Daging segar ikan tenggiri',5),
(12,N'Kepiting Segar',N'Daging segar kepiting',9),
(13,N'Mutiara',N'Mutiara yang berasal dari kerang mutiara',7),
(14,N'Makarel Kaleng',N'Daging makarel dalam kaleng',4);
INSERT INTO SUPPLIER (KdSupplier,Nama,NoTelp,Lokasi) VALUES
(1,N'Agro Panca Indonesia',N'087041086131',N'Malang'),
(2,N'Tisbox',N'080592800307',N'Malang'),
(3,N'Aquafarm Nusantara',N'088882831137',N'Pasuruan'),
(4,N'Ibu Weinny',N'087270023234',N'Pasuruan'),
(5,N'Silly Fish Indonesia',N'084398226890',N'Banyuwangi'),
(6,N'Indah Seafood',N'088763600705',N'Banyuwangi'),
(7,N'eFisheryFresh',N'089435875500',N'Surabaya'),
(8,N'Jatiluhur Mas',N'080161639494',N'Surabaya'),
(9,N'Ibu Sumartini',N'085869226315',N'Sidoarjo'),
(10,N'Fish Village',N'086548281804',N'Sidoarjo');
INSERT INTO SUPPLIER_MEMBELI_PRODUK_DARI_PABRIK (KdSupplier,KdProduk,KdPabrik,Tanggal,Jumlah) VALUES
(6,3,3,'2021-08-10',50),
(6,9,10,'2021-08-15',200),
(7,2,6,'2021-08-11',80),
(7,14,11,'2021-08-12',250),
(8,8,1,'2021-08-15',300),
(8,12,8,'2021-08-16',320);
INSERT INTO SUPPLIER_MENSUPLAI_PRODUK (KdSupplier,KdProduk,BiayaOperasional,Keuntungan,Stok) VALUES
(1,5,500,500,250),
(1,6,400,400,400),
(2,7,500,600,350),
(2,8,1000,600,200),
(3,1,400,500,500),
(3,2,1000,500,300),
(3,3,800,500,600),
(3,4,100,500,1000),
(3,6,300,500,700),
(4,8,750,750,150),
(4,9,50,50,500),
(4,10,1000,1000,100),
(4,12,1000,1000,125),
(4,14,250,250,350),
(5,2,1500,1000,400),
(5,5,500,250,600),
(5,7,600,500,650),
(5,10,1000,800,500),
(6,3,800,600,300),
(6,6,300,500,500),
(6,9,25,75,800),
(6,11,400,100,450),
(7,2,1000,500,150),
(7,7,300,200,300),
(7,9,10,40,850),
(7,14,150,50,500),
(8,1,250,150,450),
(8,4,100,100,400),
(8,8,600,150,300),
(8,12,650,50,350),
(8,13,2500,2500,50),
(9,1,300,300,500),
(9,2,1500,1500,300),
(9,3,1000,1000,350),
(10,4,250,250,700),
(10,13,6000,1500,400),
(10,14,300,200,750);

EXEC sp_MSforeachtable "ALTER TABLE ? WITH CHECK CHECK CONSTRAINT ALL"; -- Enable FK checks

