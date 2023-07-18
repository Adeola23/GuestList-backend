DROP TABLE IF EXISTS tables;
DROP TABLE IF EXISTS guest_list;



CREATE TABLE `tables` (
  `table_id` INT NOT NULL auto_increment,
  `capacity` INT NOT NULL,
  PRIMARY KEY (`table_id`)
);


CREATE TABLE `guest_list`(
  `guest_id` INT NOT NULL auto_increment,
  `guest_name` VARCHAR (50) UNIQUE NOT NULL,
  `table_id` INT NOT NULL,
  `status` VARCHAR(20) NOT NULL DEFAULT 'NOT ARRIVED',
  `accompanying_guests` INT NOT NULL,
  `time_arrived` DATETIME ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`guest_id`),
   FOREIGN KEY (`table_id`) REFERENCES tables(`table_id`)
);


