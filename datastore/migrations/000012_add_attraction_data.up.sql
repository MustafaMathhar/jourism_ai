ALTER TABLE attractions
ADD COLUMN banner_url varchar(1048) default null;
UPDATE attractions
  SET banner_url="https://competition-bucket23.s3.eu-west-3.amazonaws.com/petra_banner.jpg"
  WHERE id=1;
INSERT INTO attractions (
  name,
  description,
  lat,
  lng,mobile_phone,land_line,price,category_id,profile_id,banner_url)
VALUES
  ('متحف الأردن','يهدف متحف الأردن الى الحفاظ على الموروث الثقافي الغني للبلاد وتقديمه بطرق تفاعلية جاذبة لزائريه. يعتبر المتحف بوابة سياحية ومركزا تعليميا للتعريف عن تاريخ البلاد العريق. يحرص متحف الاردن على تقديم معلومات معرفية قيد التطوير والبحث المستمرين لكل قسم يقدمه, من قاعات العرض الى البرامج التعليمية وبرامج حفظ التراث.

يروي المتحف حكاية مليون ونصف المليون سنة من التواجد البشري والارث التاريخي على ارض الأردن, من العصر الحجري القديم بداية حتى الزمن الحاضر من خلال عرضها في قاعات التسلسل التاريخي.

يعرض المتحف مقتنياته الاثرية الثمينة بطرق فنية وتفاعلية متميزة. من تلك المقتنيات تمثال عين غزال الذي يعتبر من اقدم التماثيل البشرية في العالم, ومخطوطات البحر الميت الشهيرة وغيرها الكثير من المكتشفات التاريخية القيمة.',
  31.94554019388526, 35.927538650969076,'+96264629317',null,null,null,null,'https://competition-bucket23.s3.eu-west-3.amazonaws.com/5a401665-854b-4be3-a218-a2f6c9b186b3.jpg')
