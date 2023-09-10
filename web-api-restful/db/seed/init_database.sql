INSERT INTO public."permission" (name,code,created_at) VALUES
	 ('VIEW','VIEW','0001-01-01 07:06:30+07:06:30'),
	 ('UPDATE','UPDATE','0001-01-01 07:06:30+07:06:30'),
	 ('CREATE','CREATE','0001-01-01 07:06:30+07:06:30'),
	 ('DELETE','DELETE','0001-01-01 07:06:30+07:06:30');

INSERT INTO public."role" (name) VALUES
	 ('USER'),
	 ('MOD'),
	 ('ADMIN');

INSERT INTO public.role_permission (role_id, permission_id) VALUES
	 (3, 1),
	 (3, 2),
	 (3, 3),
	 (3, 4),
	 (1, 1);

-- {
--     "password": "123a@123",
--     "email": "liam_vo@gmail.com"
-- }
INSERT INTO public.user (username,email,phone_number,"password",created_at,updated_at) VALUES
	 ('liamvo','liam_vo@gmail.com','0989838891','$2a$14$zIQJN7AZ0G9GgZdZbsbLAenyl7FedPC0nZcIHSVQ15uhMMcbRw97C','2022-12-28 15:59:51.794425+07','2022-12-28 15:59:51.794425+07');

INSERT INTO public.user_role (user_id, role_id) VALUES
	 (1, 3);
