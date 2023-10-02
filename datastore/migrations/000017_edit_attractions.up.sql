UPDATE attractions
  SET location = CASE id
                 WHEN 1 THEN 'وادي موسى، الأردن' 
when 2 THEN 'عمان، الأردن'
  when 3 then 'الطيبية، عمان، الأردن'
  when 4 then 'عمان، الأردن' 
  else location
end
  WHERE id in(1,2,3,4);
