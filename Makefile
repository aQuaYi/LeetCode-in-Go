run:
	@echo "准备下载JSON文件，请确保浏览器已经登录了LeetCode.com"
	@wget -O algorithms.json https://leetcode.com/api/problems/algorithms/  
	@wget -O shell.json https://leetcode.com/api/problems/shell/  
	@wget -O draft.json https://leetcode.com/api/problems/draft/  
	@wget -O database.json https://leetcode.com/api/problems/database/  
	@wget -O system-design.json https://leetcode.com/api/problems/system-design/  
	@wget -O operating-system.json https://leetcode.com/api/problems/operating-system/  
	@echo "JSON文件已经下载完成了"