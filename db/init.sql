CREATE TABLE Doctors (
    First_Name varchar,
    Last_Name varchar,
    Doctor_ID varchar primary key,
    Doc_Password varchar NOT NULL
    );

CREATE TABLE Pharmacists (
    First_Name varchar,
    Last_Name varchar,
    Employee_ID varchar primary key,
    Pharm_Password varchar NOT NULL,
    Is_Manager boolean
        );

CREATE TABLE Inventory (
    Drug_Name varchar primary key,
    Amt_On_Hand int,
    Cost_Per_Mg decimal,
    Supplier varchar
);

CREATE TABLE Prescriptions (
    Presc_ID varchar,
    Doc_Name varchar,
    Drug_Name varchar,
    Amount int,
    Patient_First varchar,
    Patient_Last varchar,
    Cost decimal,
    Presc_Status varchar,
    Date_Prescribed date
);

CREATE TABLE Prescription_History (
    Presc_ID varchar,
    Doc_Prescribing varchar,
    Drug_Name varchar,
    Amount int,
    Patient_First varchar,
    Patient_Last varchar,
    Cost decimal,
    Date_Prescribed date
);

INSERT INTO Doctors values ('Young', 'Farwa', '123456', 'narwall');
INSERT INTO Pharmacists values ('Bruce', 'Banner', '789101', 'hulksmash','true');
INSERT INTO Inventory values ('Ibuprofen', '005500', '1.25', 'Meditech');
INSERT INTO Prescriptions values ('459056', 'Young Farwa', 'Amoxicillin', '500', 'Tony', 'Stark', '31.29', 'filled', '2018-10-31');
INSERT INTO Prescription_History values ('459055', 'Young Farwa', 'Amoxicillin', '500', 'Tony', 'Stark', '31.29', 'picked-up', '2018-10-31');
