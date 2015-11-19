# YoutubeDownloadServer  

1.先依作業系統安裝好ffmpeg、youtube-dl套件    

2.執行(未指定port,預設7777)   
```
./youtube-server 8080  
```

3.預設帳號:go、密碼:go  
4.其它參數說明  
單一檔:  
127.0.0.1:8080/video?v=qWnKYCjANP8     -> 捉video  
127.0.0.1:8080/encode??v=qWnKYCjANP8   -> 轉成mp3  
127.0.0.1:8080/download web目錄         -> 捉完直接看一下有那些檔案  

直接捉清單所有檔案:  
127.0.0.1:8080/video?list=PLgrxkk47XLR5K6lU3m3-6pjfodz_z8tPQ  
127.0.0.1:8080/encode?list=PLgrxkk47XLR5K6lU3m3-6pjfodz_z8tPQ  

<img src="https://lh3.googleusercontent.com/-G_Zz8kWXzvM/Vk4OweZSXbI/AAAAAAAAO8U/TP0y9bAhC78/s720-Ic42/%2525E8%25259E%2525A2%2525E5%2525B9%252595%2525E5%2525BF%2525AB%2525E7%252585%2525A7%2525202015-11-20%252520%2525E4%2525B8%25258A%2525E5%25258D%2525881.58.19.png"
alt="IMAGE ALT TEXT HERE" width="640" height="280" border="10" />
<img src="https://lh3.googleusercontent.com/-R747USLOAjc/Vk4OwWlJm_I/AAAAAAAAO8Y/NJvfCzWiooM/s576-Ic42/%2525E8%25259E%2525A2%2525E5%2525B9%252595%2525E5%2525BF%2525AB%2525E7%252585%2525A7%2525202015-11-20%252520%2525E4%2525B8%25258A%2525E5%25258D%2525881.58.41.png"
alt="IMAGE ALT TEXT HERE" width="640" height="280" border="10" />
