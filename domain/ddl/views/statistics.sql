CREATE VIEW statistics
AS
SELECT 
url_id,
COUNT(CASE WHEN access_dt < date('now','-24 hours') THEN 1 END) AS twenty_four_hours,
COUNT(CASE WHEN access_dt < date('now','-7 days') THEN 1 END) AS last_seven_days,
COUNT(*) AS all_time
FROM HISTORY
GROUP BY url_id