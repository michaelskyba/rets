CREATE TABLE residential_records (
	id			INT AUTO_INCREMENT NOT NULL,
	rets_date	DATE NOT NULL,

	A_c VARCHAR(11),
	Addr VARCHAR(35),
	Ad_text VARCHAR(463),
	All_inc VARCHAR(1),
	Area VARCHAR(40),
	Area_code VARCHAR(4),
	Bath_tot INT,
	Br INT,
	Bsmt1_out VARCHAR(12),
	Cable VARCHAR(1),
	Cac_inc VARCHAR(1),
	Com_coopb VARCHAR(40),
	Comel_inc VARCHAR(1),
	Community VARCHAR(44),
	Community_code VARCHAR(20),
	Comp_pts VARCHAR(1),
	Constr1_out VARCHAR(16),
	County VARCHAR(16),
	Cross_st VARCHAR(30),
	Den_fr VARCHAR(1),
	Depth INT,
	Disp_addr VARCHAR(1),
	Dom INT,
	Drive VARCHAR(10),
	Extras VARCHAR(240),
	Fpl_num VARCHAR(1),
	Front_ft INT,
	Fuel VARCHAR(9),
	Furnished VARCHAR(4),
	Gar_spaces INT,
	Gar_type VARCHAR(16),
	Heat_inc VARCHAR(1),
	Heating VARCHAR(20),
	Hydro_inc VARCHAR(1),
	Input_date DATE,
	Internet VARCHAR(1),
	Laundry VARCHAR(13),
	Laundry_lev VARCHAR(5),
	Ld DATE,
	Lease VARCHAR(1),
	Lease_term VARCHAR(10),
	Legal_desc VARCHAR(50),
	Level1 VARCHAR(8),
	Level2 VARCHAR(8),
	Level3 VARCHAR(8),
	Level4 VARCHAR(8),
	Level5 VARCHAR(8),
	Level8 VARCHAR(8),
	Link_yn VARCHAR(1),
	Lotsz_code VARCHAR(8),
	Lp_dol INT,
	Lsc VARCHAR(3),
	Ml_num VARCHAR(8),
	Mmap_col INT,
	Mmap_page INT,
	Mmap_row VARCHAR(1),
	Municipality VARCHAR(40),
	Municipality_code VARCHAR(10),
	Municipality_district VARCHAR(44),
	Num_kit INT,
	Occ VARCHAR(14),
	Orig_dol INT,
	Park_spcs INT,
	Pix_updt DATE,
	Pool VARCHAR(8),
	Portion_property_lease_srch VARCHAR(50),
	Poss_date DATE,
	Prkg_inc VARCHAR(1),
	Pvt_ent VARCHAR(1),
	Rltr VARCHAR(80),
	Rm1_dc1_out VARCHAR(20),
	Rm1_len INT,
	Rm1_out VARCHAR(9),
	Rm1_wth INT,
	Rm2_dc1_out VARCHAR(20),
	Rm2_len INT,
	Rm2_out VARCHAR(9),
	Rm2_wth INT,
	Rm3_dc1_out VARCHAR(20),
	Rm3_len INT,
	Rm3_out VARCHAR(9),
	Rm3_wth INT,
	Rm4_dc1_out VARCHAR(20),
	Rm6_out VARCHAR(9),
	Rms INT,
	Rooms_plus INT,
	Sewer VARCHAR(15),
	Spec_des1_out VARCHAR(27),
	Sqft VARCHAR(9),
	S_r VARCHAR(9),
	St VARCHAR(20),
	Status VARCHAR(1),
	St_num VARCHAR(7),
	St_sfx VARCHAR(4),
	Style VARCHAR(16),
	Taxes INT,
	Timestamp_sql DATE,
	Tot_park_spcs INT,
	Tour_url VARCHAR(100),
	Type_own1_out VARCHAR(18),
	Type_own_srch VARCHAR(7),
	Vend_pis VARCHAR(10),
	Water VARCHAR(9),
	Water_inc VARCHAR(1),
	Wcloset_p1 INT,
	Wcloset_p2 INT,
	Wcloset_t1 INT,
	Wcloset_t1lvl VARCHAR(8),
	Wcloset_t2 INT,
	Wcloset_t2lvl VARCHAR(8),
	Xd DATE,
	Yr INT,
	Yr_built VARCHAR(5),
	Zip VARCHAR(7),

	PRIMARY KEY (`id`)
);
