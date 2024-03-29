CREATE TABLE Doctors (
    Doc_Username varchar primary key,
    First_Name varchar,
    Last_Name varchar,
    Doc_Password varchar);

CREATE TABLE Pharmacists (
    First_Name varchar,
    Last_Name varchar,
    Username varchar primary key,
    Pharm_Password varchar,
    Is_Manager varchar);

CREATE TABLE Inventory (
    Drug_Name varchar primary key,
    Amt_On_Hand int,
    Cost_Per_Mg decimal,
    Supplier varchar);

CREATE TABLE Prescriptions (
    Presc_ID varchar,
    Doc_Name varchar,
    Drug_Name varchar,
    Amount int,
    Patient_First varchar,
    Patient_Last varchar,
    Cost decimal,
    Presc_Status varchar,
    Date_Prescribed varchar
);

CREATE TABLE Prescription_History (
    Presc_ID varchar,
    Doc_Prescribing varchar,
    Drug_Name varchar,
    Amount int,
    Patient_First varchar,
    Patient_Last varchar,
    Cost decimal,
    Date_Prescribed varchar);

INSERT INTO Doctors
values ('drFarwa', 'Young', 'Farwa', 'thefarwacist');

INSERT INTO Doctors
values ('drStrange', 'Strange', 'stephen', 'timeStone');

INSERT INTO Pharmacists 
values ('Bruce', 'Banner', 'MrGreen', 'hulksmash','true');

INSERT INTO Inventory 
values ('Ibuprofen', 5500, 1.25, 'Meditech');

INSERT INTO Prescriptions 
values ('459056', 'drFarwa', 'Amoxicillin', 500, 'Tony', 'Stark', 31.25, 'filled', '2018-10-31');

INSERT INTO Prescriptions 
values ('318275', 'drFarwa', 'Synthroid', 500, 'Scarlet', 'witch', 31.25, 'filled', '2018-3-14');

INSERT INTO Prescriptions 
values ('563527', 'drStrange', 'Nexium', 500, 'Thor', 'Odinson', 31.25, 'filled', '2018-12-24');

INSERT INTO Prescription_History 
values ('459056', 'drFarwa', 'Amoxicillin', 500, 'Tony', 'Stark', 31.25, '2018-10-31');

INSERT INTO Prescription_History
values ('318275', 'drFarwa', 'Synthroid', 500, 'Scarlet', 'witch', 31.25, '2018-3-14');

INSERT INTO Prescription_History
values ('563527', 'drStrange', 'Nexium', 500, 'Thor', 'Odinson', 31.25, '2018-12-24');
