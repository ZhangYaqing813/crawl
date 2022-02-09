package controllers

import (
	"crawl_books/models"
	"fmt"
	"github.com/astaxie/beego"
)

type BooksController struct {
	beego.Controller
}

func (c *BooksController) BooksInfo() {

	bookHtml := `

<!DOCTYPE html>
<html lang="zh-cmn-Hans" class="ua-windows ua-webkit book-new-nav">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <title>
  豆瓣图书标签: 编程
</title>
  
<script>!function(e){var o=function(o,n,t){var c,i,r=new Date;n=n||30,t=t||"/",r.setTime(r.getTime()+24*n*60*60*1e3),c="; expires="+r.toGMTString();for(i in o)e.cookie=i+"="+o[i]+c+"; path="+t},n=function(o){var n,t,c,i=o+"=",r=e.cookie.split(";");for(t=0,c=r.length;t<c;t++)if(n=r[t].replace(/^\s+|\s+$/g,""),0==n.indexOf(i))return n.substring(i.length,n.length).replace(/\"/g,"");return null},t=e.write,c={"douban.com":1,"douban.fm":1,"google.com":1,"google.cn":1,"googleapis.com":1,"gmaptiles.co.kr":1,"gstatic.com":1,"gstatic.cn":1,"google-analytics.com":1,"googleadservices.com":1},i=function(e,o){var n=new Image;n.onload=function(){},n.src="https://www.douban.com/j/except_report?kind=ra022&reason="+encodeURIComponent(e)+"&environment="+encodeURIComponent(o)},r=function(o){try{t.call(e,o)}catch(e){t(o)}},a=/<script.*?src\=["']?([^"'\s>]+)/gi,g=/http:\/\/(.+?)\.([^\/]+).+/i;e.writeln=e.write=function(e){var t,l=a.exec(e);return l&&(t=g.exec(l[1]))?c[t[2]]?void r(e):void("tqs"!==n("hj")&&(i(l[1],location.href),o({hj:"tqs"},1),setTimeout(function(){location.replace(location.href)},50))):void r(e)}}(document);</script>

  
  <meta http-equiv="Pragma" content="no-cache">
  <meta http-equiv="Expires" content="Sun, 6 Mar 2005 01:00:00 GMT">
  
  <script>var _head_start = new Date();</script>
  
  <link href="https://img3.doubanio.com/f/book/8011541306494bc2ff483299d952d9c65753c43d/css/book/master.css" rel="stylesheet" type="text/css">

  <link href="https://img3.doubanio.com/f/book/222a5c61e041638af8defc87cf97f4a863a77922/css/book/base/init.css" rel="stylesheet">
  <style type="text/css"></style>
  <script src="https://img3.doubanio.com/f/book/0495cb173e298c28593766009c7b0a953246c5b5/js/book/lib/jquery/jquery.js"></script>
  <script src="https://img3.doubanio.com/f/shire/22ee83f45f94c7a90e73e0ee4acd18f902a6991f/js/douban.js"></script>
  <script src="https://img3.doubanio.com/f/book/0322e3e810e475f1c82adb7d1c6ccfa1c0fa969c/js/book/master.js"></script>
  

  
  <script>  </script>
  <link rel="stylesheet" href="https://img3.doubanio.com/misc/mixed_static/468bc24284a3074.css">

  <link rel="shortcut icon" href="https://img3.doubanio.com/favicon.ico" type="image/x-icon">
</head>
<body>
  
    <script>var _body_start = new Date();</script>
    
  



    <link href="//img3.doubanio.com/dae/accounts/resources/3e96b44/shire/bundle.css" rel="stylesheet" type="text/css">



<div id="db-global-nav" class="global-nav">
  <div class="bd">
    
<div class="top-nav-info">
  <a href="https://accounts.douban.com/passport/login?source=book" class="nav-login" rel="nofollow">登录/注册</a>
</div>


    <div class="top-nav-doubanapp">
  <a href="https://www.douban.com/doubanapp/app?channel=top-nav" class="lnk-doubanapp">下载豆瓣客户端</a>
  <div id="doubanapp-tip">
    <a href="https://www.douban.com/doubanapp/app?channel=qipao" class="tip-link">豆瓣 <span class="version">6.0</span> 全新发布</a>
    <a href="javascript: void 0;" class="tip-close">×</a>
  </div>
  <div id="top-nav-appintro" class="more-items">
    <p class="appintro-title">豆瓣</p>
    <p class="qrcode">扫码直接下载</p>
    <div class="download">
      <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=iOS">iPhone</a>
      <span>·</span>
      <a href="https://www.douban.com/doubanapp/redirect?channel=top-nav&direct_dl=1&download=Android" class="download-android">Android</a>
    </div>
  </div>
</div>

    


<div class="global-nav-items">
  <ul>
    <li class="">
      <a href="https://www.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-main&quot;,&quot;uid&quot;:&quot;0&quot;}">豆瓣</a>
    </li>
    <li class="on">
      <a href="https://book.douban.com"  data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-book&quot;,&quot;uid&quot;:&quot;0&quot;}">读书</a>
    </li>
    <li class="">
      <a href="https://movie.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-movie&quot;,&quot;uid&quot;:&quot;0&quot;}">电影</a>
    </li>
    <li class="">
      <a href="https://music.douban.com" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-music&quot;,&quot;uid&quot;:&quot;0&quot;}">音乐</a>
    </li>
    <li class="">
      <a href="https://www.douban.com/location" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-location&quot;,&quot;uid&quot;:&quot;0&quot;}">同城</a>
    </li>
    <li class="">
      <a href="https://www.douban.com/group" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-group&quot;,&quot;uid&quot;:&quot;0&quot;}">小组</a>
    </li>
    <li class="">
      <a href="https://read.douban.com&#47;?dcs=top-nav&amp;dcm=douban" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-read&quot;,&quot;uid&quot;:&quot;0&quot;}">阅读</a>
    </li>
    <li class="">
      <a href="https://douban.fm&#47;?from_=shire_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-fm&quot;,&quot;uid&quot;:&quot;0&quot;}">FM</a>
    </li>
    <li class="">
      <a href="https://time.douban.com&#47;?dt_time_source=douban-web_top_nav" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-time&quot;,&quot;uid&quot;:&quot;0&quot;}">时间</a>
    </li>
    <li class="">
      <a href="https://market.douban.com&#47;?utm_campaign=douban_top_nav&amp;utm_source=douban&amp;utm_medium=pc_web" target="_blank" data-moreurl-dict="{&quot;from&quot;:&quot;top-nav-click-market&quot;,&quot;uid&quot;:&quot;0&quot;}">豆品</a>
    </li>
  </ul>
</div>

  </div>
</div>
<script>
  ;window._GLOBAL_NAV = {
    DOUBAN_URL: "https://www.douban.com",
    N_NEW_NOTIS: 0,
    N_NEW_DOUMAIL: 0
  };
</script>



    <script src="//img3.doubanio.com/dae/accounts/resources/3e96b44/shire/bundle.js" defer="defer"></script>




  



    <link href="//img3.doubanio.com/dae/accounts/resources/3e96b44/book/bundle.css" rel="stylesheet" type="text/css">




<div id="db-nav-book" class="nav">
  <div class="nav-wrap">
  <div class="nav-primary">
    <div class="nav-logo">
      <a href="https:&#47;&#47;book.douban.com">豆瓣读书</a>
    </div>
    <div class="nav-search">
      <form action="https:&#47;&#47;search.douban.com&#47;book/subject_search" method="get">
        <fieldset>
          <legend>搜索：</legend>
          <label for="inp-query">
          </label>
          <div class="inp"><input id="inp-query" name="search_text" size="22" maxlength="60" placeholder="书名、作者、ISBN" value=""></div>
          <div class="inp-btn"><input type="submit" value="搜索"></div>
          <input type="hidden" name="cat" value="1001" />
        </fieldset>
      </form>
    </div>
  </div>
  </div>
  <div class="nav-secondary">
    

<div class="nav-items">
  <ul>
    <li    ><a href="https://book.douban.com/cart/"
     >购书单</a>
    </li>
    <li    ><a href="https://read.douban.com/ebooks/?dcs=book-nav&dcm=douban"
            target="_blank"
     >电子图书</a>
    </li>
    <li    ><a href="https://market.douban.com/book?utm_campaign=book_nav_freyr&utm_source=douban&utm_medium=pc_web"
     >豆瓣书店</a>
    </li>
    <li    ><a href="https://book.douban.com/annual/2021?source=navigation"
            target="_blank"
     >2021年度榜单</a>
    </li>
    <li    ><a href="https://www.douban.com/standbyme/2021?fullscreen=true&hidenav=true&autorotate=false&source=book_navigation"
            target="_blank"
     >2021书影音报告</a>
    </li>
    <li          class=" book-cart"
    ><a href="https://market.douban.com/cart/?biz_type=book&utm_campaign=book_nav_cart&utm_source=douban&utm_medium=pc_web"
            target="_blank"
     >购物车</a>
    </li>
  </ul>
</div>

    <a href="https://book.douban.com/annual/2021?source=book_navigation" class="bookannual"></a>
  </div>
</div>

<script id="suggResult" type="text/x-jquery-tmpl">
  <li data-link="{{= url}}">
            <a href="{{= url}}" onclick="moreurl(this, {from:'book_search_sugg', query:'{{= keyword }}', subject_id:'{{= id}}', i: '{{= index}}', type: '{{= type}}'})">
            <img src="{{= pic}}" width="40" />
            <div>
                <em>{{= title}}</em>
                {{if year}}
                    <span>{{= year}}</span>
                {{/if}}
                <p>
                {{if type == "b"}}
                    {{= author_name}}
                {{else type == "a" }}
                    {{if en_name}}
                        {{= en_name}}
                    {{/if}}
                {{/if}}
                 </p>
            </div>
        </a>
        </li>
  </script>




    <script src="//img3.doubanio.com/dae/accounts/resources/3e96b44/book/bundle.js" defer="defer"></script>





    <div id="wrapper">
        
        
  <div id="content">
    
    <h1>豆瓣图书标签: 编程</h1>

    <div class="grid-16-8 clearfix">
      
      <div class="article">
  <div id="subject_list">

      
  
  
  <div class="clearfix">
    <span class="rr greyinput">

          综合排序
          &nbsp;/&nbsp;

          <a href="/tag/%E7%BC%96%E7%A8%8B?type=R">按出版日期排序</a>
          &nbsp;/&nbsp;

          <a href="/tag/%E7%BC%96%E7%A8%8B?type=S">按评价排序</a>
    </span>
  </div>



    
  
  <ul class="subject-list">
      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/35196328/" 
  onclick="moreurl(this,{i:'0',query:'',subject_id:'35196328',from:'book_subject_search'})">
        <img class="" src="https://img1.doubanio.com/view/subject/s/public/s33716278.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/35196328/" title="Python编程" 
  onclick="moreurl(this,{i:'0',query:'',subject_id:'35196328',from:'book_subject_search'})">

    Python编程


    
      <span style="font-size:12px;"> : 从入门到实践（第2版） </span>

  </a>

      </h2>
      <div class="pub">
        
  
  [美]埃里克·马瑟斯（Eric Matthes） / 袁国忠 / 人民邮电出版社 / 2020-10 / 89

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.2</span>

    <span class="pl">
        (722人评价)
    </span>
  </div>



        
  
  
  
    <p>本书是针对所有层次Python读者而作的Python入门书。全书分两部分：第一部分介绍用Python编程所必须了解的基本概念，包括Matplotlib等强大... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/35196328/buylinks">
        纸质版 109.80元
      </a>
    </span>

          </div>

            
            
  
    
    
  <div class="ebook-link">
    <a target="_blank" href="https://read.douban.com/ebook/337008762/?dcs=tag-buylink&amp;dcm=douban&amp;dct=35196328">去看电子版</a>
  </div>
  


      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/35641088/" 
  onclick="moreurl(this,{i:'1',query:'',subject_id:'35641088',from:'book_subject_search'})">
        <img class="" src="https://img3.doubanio.com/view/subject/s/public/s34032260.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/35641088/" title="计算之魂" 
  onclick="moreurl(this,{i:'1',query:'',subject_id:'35641088',from:'book_subject_search'})">

    计算之魂


    
      <span style="font-size:12px;"> : 计算科学品位和认知进阶 </span>

  </a>

      </h2>
      <div class="pub">
        
  
  吴军 / 人民邮电出版社 / 2021-11 / 109

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.3</span>

    <span class="pl">
        (79人评价)
    </span>
  </div>



        
  
  
  
    <p>对计算机科学的掌握程度，决定了一个计算机行业从业者能走多远。在本书中，作者将人文历史与计算机科学相结合，通过一些具体的例题，分10个主题系统地讲解了计算机科... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/35641088/buylinks">
        纸质版 85.90元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/35524665/" 
  onclick="moreurl(this,{i:'2',query:'',subject_id:'35524665',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s34038775.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/35524665/" title="JavaScript百炼成仙" 
  onclick="moreurl(this,{i:'2',query:'',subject_id:'35524665',from:'book_subject_search'})">

    JavaScript百炼成仙


    

  </a>

      </h2>
      <div class="pub">
        
  
  杨逸飞 / 2021-6-1 / 66.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar40"></span>
        <span class="rating_nums">7.7</span>

    <span class="pl">
        (61人评价)
    </span>
  </div>



        
  
  
  
    <p>这是一本讲解JavaScript编程语言的技术书籍，只不过，本书采用了一种全新的写作手法。如果你厌倦了厚厚的、如同字典般的编程书籍，不妨尝试一下新的口味！通... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/35524665/buylinks">
        纸质版 56.10元
      </a>
    </span>

          </div>

            
            
  
    
    
  <div class="ebook-link">
    <a target="_blank" href="https://read.douban.com/ebook/326306541/?dcs=tag-buylink&amp;dcm=douban&amp;dct=35524665">去看电子版</a>
  </div>
  


      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/4822685/" 
  onclick="moreurl(this,{i:'3',query:'',subject_id:'4822685',from:'book_subject_search'})">
        <img class="" src="https://img2.doubanio.com/view/subject/s/public/s27331702.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/4822685/" title="编码" 
  onclick="moreurl(this,{i:'3',query:'',subject_id:'4822685',from:'book_subject_search'})">

    编码


    
      <span style="font-size:12px;"> : 隐匿在计算机软硬件背后的语言 </span>

  </a>

      </h2>
      <div class="pub">
        
  
  [美] Charles Petzold / 左飞、薛佟佟 / 电子工业出版社 / 2010 / 55.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.3</span>

    <span class="pl">
        (4406人评价)
    </span>
  </div>



        
  
  
  
    <p>本书讲述的是计算机工作原理。作者用丰富的想象和清晰的笔墨将看似繁杂的理论阐述得通俗易懂，你丝毫不会感到枯燥和生硬。更重要的是，你会因此而获得对计算机工作原理... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/4822685/buylinks">
        纸质版 46.10元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/30329536/" 
  onclick="moreurl(this,{i:'4',query:'',subject_id:'30329536',from:'book_subject_search'})">
        <img class="" src="https://img2.doubanio.com/view/subject/s/public/s29872642.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/30329536/" title="数据密集型应用系统设计" 
  onclick="moreurl(this,{i:'4',query:'',subject_id:'30329536',from:'book_subject_search'})">

    数据密集型应用系统设计


    

  </a>

      </h2>
      <div class="pub">
        
  
  Martin Kleppmann / 赵军平、李三平、吕云松、耿煜 / 中国电力出版社 / 2018-9-1 / 128

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar50"></span>
        <span class="rating_nums">9.7</span>

    <span class="pl">
        (1085人评价)
    </span>
  </div>



        
  
  
  
    <p>全书分为三大部分：
第一部分，主要讨论有关增强数据密集型应用系统所需的若干基本原则。首先开篇第1章即瞄准目标：可靠性、可扩展性与可维护性，如何认识这些问题以... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/30329536/buylinks">
        纸质版 113.80元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/6021440/" 
  onclick="moreurl(this,{i:'5',query:'',subject_id:'6021440',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s4669554.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/6021440/" title="黑客与画家" 
  onclick="moreurl(this,{i:'5',query:'',subject_id:'6021440',from:'book_subject_search'})">

    黑客与画家


    
      <span style="font-size:12px;"> : 硅谷创业之父Paul Graham文集 </span>

  </a>

      </h2>
      <div class="pub">
        
  
  [美] Paul Graham / 阮一峰 / 人民邮电出版社 / 2011-4 / 49.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">8.7</span>

    <span class="pl">
        (21276人评价)
    </span>
  </div>



        
  
  
  
    <p>本书是硅谷创业之父Paul Graham 的文集，主要介绍黑客即优秀程序员的爱好和动机，讨论黑客成长、黑客对世界的贡献以及编程语言和黑客工作方法等所有对计算... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/6021440/buylinks">
        纸质版 60.70元
      </a>
    </span>

          </div>

            
            
  
    
    
  <div class="ebook-link">
    <a target="_blank" href="https://read.douban.com/ebook/163682605/?dcs=tag-buylink&amp;dcm=douban&amp;dct=6021440">去看电子版</a>
  </div>
  


      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/26912767/" 
  onclick="moreurl(this,{i:'6',query:'',subject_id:'26912767',from:'book_subject_search'})">
        <img class="" src="https://img1.doubanio.com/view/subject/s/public/s29195878.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/26912767/" title="深入理解计算机系统（原书第3版）" 
  onclick="moreurl(this,{i:'6',query:'',subject_id:'26912767',from:'book_subject_search'})">

    深入理解计算机系统（原书第3版）


    

  </a>

      </h2>
      <div class="pub">
        
  
  Randal E.Bryant、David O&#39;Hallaron / 龚奕利、贺莲 / 机械工业出版社 / 2016-11 / 139.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar50"></span>
        <span class="rating_nums">9.8</span>

    <span class="pl">
        (1718人评价)
    </span>
  </div>



        
  
  
  
    <p>和第2版相比，本版内容上*大的变化是，从以IA32和x86-64为基础转变为完全以x86-64为基础。主要更新如下：
基于x86-64，大量地重写代码，首次... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
              <span class="market-info">
                <a href="https://book.douban.com/subject/26912767/?channel=subject_list&amp;platform=web" target="_blank">在豆瓣购买</a>
              </span>
          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/1148282/" 
  onclick="moreurl(this,{i:'7',query:'',subject_id:'1148282',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s1113106.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/1148282/" title="计算机程序的构造和解释(原书第2版)" 
  onclick="moreurl(this,{i:'7',query:'',subject_id:'1148282',from:'book_subject_search'})">

    计算机程序的构造和解释(原书第2版)


    

  </a>

      </h2>
      <div class="pub">
        
  
  [美] Harold Abelson、[美] Gerald Jay Sussman、[美] Julie Sussman / 裘宗燕 / 机械工业出版社 / 2004-2 / 45.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar50"></span>
        <span class="rating_nums">9.5</span>

    <span class="pl">
        (2542人评价)
    </span>
  </div>



        
  
  
  
    <p>《计算机程序的构造和解释》成型于美国麻省理工学院（MIT）多年使用的一本教材，1984年出版，1996年修订为第二版。在过去的二十多年里，该书对于计算机科学... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/1148282/buylinks">
        纸质版 34.90元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/5333562/" 
  onclick="moreurl(this,{i:'8',query:'',subject_id:'5333562',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s4510534.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/5333562/" title="深入理解计算机系统（原书第2版）" 
  onclick="moreurl(this,{i:'8',query:'',subject_id:'5333562',from:'book_subject_search'})">

    深入理解计算机系统（原书第2版）


    

  </a>

      </h2>
      <div class="pub">
        
  
  [美] Randal E.Bryant、[美] David O&#39; Hallaron / 龚奕利、雷迎春 / 机械工业出版社 / 2011-1-1 / 99.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar50"></span>
        <span class="rating_nums">9.7</span>

    <span class="pl">
        (2928人评价)
    </span>
  </div>



        
  
  
  
    <p>本书从程序员的视角详细阐述计算机系统的本质概念，并展示这些概念如何实实在在地影响应用程序的正确性、性能和实用性。全书共12章，主要内容包括信息的表示和处理、... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/5333562/buylinks">
        纸质版 81.70元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/30293801/" 
  onclick="moreurl(this,{i:'9',query:'',subject_id:'30293801',from:'book_subject_search'})">
        <img class="" src="https://img1.doubanio.com/view/subject/s/public/s29839337.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/30293801/" title="Python深度学习" 
  onclick="moreurl(this,{i:'9',query:'',subject_id:'30293801',from:'book_subject_search'})">

    Python深度学习


    

  </a>

      </h2>
      <div class="pub">
        
  
  [美] 弗朗索瓦•肖莱 / 张亮 / 人民邮电出版社 / 2018-8 / 119.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar50"></span>
        <span class="rating_nums">9.5</span>

    <span class="pl">
        (893人评价)
    </span>
  </div>



        
  
  
  
    <p>本书由Keras之父、现任Google人工智能研究员的弗朗索瓦•肖莱（François Chollet）执笔，详尽介绍了用Python和Keras进行深度学... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/30293801/buylinks">
        纸质版 99.70元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/35006892/" 
  onclick="moreurl(this,{i:'10',query:'',subject_id:'35006892',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s33836286.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/35006892/" title="程序员修炼之道（第2版）" 
  onclick="moreurl(this,{i:'10',query:'',subject_id:'35006892',from:'book_subject_search'})">

    程序员修炼之道（第2版）


    
      <span style="font-size:12px;"> : 通向务实的最高境界 </span>

  </a>

      </h2>
      <div class="pub">
        
  
  [美] David Thomas、[美] Andrew Hunt / 云风 / 电子工业出版社 / 2020-4-1 / 89.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.1</span>

    <span class="pl">
        (483人评价)
    </span>
  </div>



        
  
  
  
    <p>本书之所以在全球范围内广泛传播，被一代代开发者奉为圭臬，盖因它可以创造出真正的价值：或编写出更好的软件，或探究出编程的本质，而所有收获均不依赖于特定语言、框... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
              <span class="market-info">
                <a href="https://book.douban.com/subject/35006892/?channel=subject_list&amp;platform=web" target="_blank">在豆瓣购买</a>
              </span>
          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/1477390/" 
  onclick="moreurl(this,{i:'11',query:'',subject_id:'1477390',from:'book_subject_search'})">
        <img class="" src="https://img1.doubanio.com/view/subject/s/public/s1495029.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/1477390/" title="代码大全（第2版）" 
  onclick="moreurl(this,{i:'11',query:'',subject_id:'1477390',from:'book_subject_search'})">

    代码大全（第2版）


    

  </a>

      </h2>
      <div class="pub">
        
  
  [美] 史蒂夫·迈克康奈尔 / 金戈、汤凌、陈硕、张菲 译、裘宗燕 审校 / 电子工业出版社 / 2006-3 / 128.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.3</span>

    <span class="pl">
        (4065人评价)
    </span>
  </div>



        
  
  
  
    <p>第2版的《代码大全》是著名IT畅销书作者史蒂夫·迈克康奈尔11年前的经典著作的全新演绎：第2版不是第一版的简单修订增补，而是完全进行了重写；增加了很多与时俱... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/1477390/buylinks">
        纸质版 82.10元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/35218199/" 
  onclick="moreurl(this,{i:'12',query:'',subject_id:'35218199',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s33732686.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/35218199/" title="机器学习实战（原书第2版）" 
  onclick="moreurl(this,{i:'12',query:'',subject_id:'35218199',from:'book_subject_search'})">

    机器学习实战（原书第2版）


    
      <span style="font-size:12px;"> : 基于Scikit-Learn、Keras和TensorFlow </span>

  </a>

      </h2>
      <div class="pub">
        
  
  [法] Aurélien Géron / 宋能辉、李娴 / 机械工业出版社 / 2020-10-1 / 149.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar50"></span>
        <span class="rating_nums">9.9</span>

    <span class="pl">
        (184人评价)
    </span>
  </div>



        
  
  
  
    <p>这本机器学习畅销书基于TensorFlow 2和Scikit-Learn的新版本进行了全面更新，通过具体的示例、非常少的理论和可用于生产环境的Python框... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/35218199/buylinks">
        纸质版 149.00元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/27028517/" 
  onclick="moreurl(this,{i:'13',query:'',subject_id:'27028517',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s29434304.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/27028517/" title="流畅的Python" 
  onclick="moreurl(this,{i:'13',query:'',subject_id:'27028517',from:'book_subject_search'})">

    流畅的Python


    

  </a>

      </h2>
      <div class="pub">
        
  
  [巴西] Luciano Ramalho / 安道、吴珂 / 人民邮电出版社 / 2017-5-15 / 139元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.4</span>

    <span class="pl">
        (975人评价)
    </span>
  </div>



        
  
  
  
    <p>【技术大咖推荐】
“很荣幸担任这本优秀图书的技术审校。这本书能帮助很多中级Python程序员掌握这门语言，我也从中学到了相当多的知识！”——Alex Mar... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/27028517/buylinks">
        纸质版 139.00元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/25756435/" 
  onclick="moreurl(this,{i:'14',query:'',subject_id:'25756435',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s27244594.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/25756435/" title="逐梦旅程：Windows游戏编程之从零开始" 
  onclick="moreurl(this,{i:'14',query:'',subject_id:'25756435',from:'book_subject_search'})">

    逐梦旅程：Windows游戏编程之从零开始


    

  </a>

      </h2>
      <div class="pub">
        
  
  毛星云 / 清华大学出版社 / 2013-9-16 / 98.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">8.7</span>

    <span class="pl">
        (64人评价)
    </span>
  </div>



        
  
  
  
    <p>端游开发是目前最热的职业，报酬丰厚且能实现自己的游戏梦想。作者历经一年时间，编写了这本详细讲解Windows游戏开发的入门图书。
《逐梦旅程：Windows... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/25756435/buylinks">
        纸质版 83.30元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/19952400/" 
  onclick="moreurl(this,{i:'15',query:'',subject_id:'19952400',from:'book_subject_search'})">
        <img class="" src="https://img2.doubanio.com/view/subject/s/public/s29107491.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/19952400/" title="算法（第4版）" 
  onclick="moreurl(this,{i:'15',query:'',subject_id:'19952400',from:'book_subject_search'})">

    算法（第4版）


    

  </a>

      </h2>
      <div class="pub">
        
  
  [美] Robert Sedgewick、[美] Kevin Wayne / 谢路云 / 人民邮电出版社 / 2012-10-1 / 99.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.4</span>

    <span class="pl">
        (1617人评价)
    </span>
  </div>



        
  
  
  
    <p>本书作为算法领域经典的参考书，全面介绍了关于算法和数据结构的必备知识，并特别针对排序、搜索、图处理和字符串处理进行了论述。第4版具体给出了每位程序员应知应会... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/19952400/buylinks">
        纸质版 129.80元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/26979890/" 
  onclick="moreurl(this,{i:'16',query:'',subject_id:'26979890',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s29358625.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/26979890/" title="算法图解" 
  onclick="moreurl(this,{i:'16',query:'',subject_id:'26979890',from:'book_subject_search'})">

    算法图解


    

  </a>

      </h2>
      <div class="pub">
        
  
  [美] Aditya Bhargava / 袁国忠 / 人民邮电出版社 / 2017-3 / 49.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">8.5</span>

    <span class="pl">
        (2528人评价)
    </span>
  </div>



        
  
  
  
    <p>本书示例丰富，图文并茂，以让人容易理解的方式阐释了算法，旨在帮助程序员在日常项目中更好地发挥算法的能量。书中的前三章将帮助你打下基础，带你学习二分查找、大O... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/26979890/buylinks">
        纸质版 38.80元
      </a>
    </span>

          </div>

            
            
  
    
    
  <div class="ebook-link">
    <a target="_blank" href="https://read.douban.com/ebook/52186827/?dcs=tag-buylink&amp;dcm=douban&amp;dct=26979890">去看电子版</a>
  </div>
  


      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/35387685/" 
  onclick="moreurl(this,{i:'17',query:'',subject_id:'35387685',from:'book_subject_search'})">
        <img class="" src="https://img9.doubanio.com/view/subject/s/public/s33845055.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/35387685/" title="Python编程快速上手（第2版）" 
  onclick="moreurl(this,{i:'17',query:'',subject_id:'35387685',from:'book_subject_search'})">

    Python编程快速上手（第2版）


    
      <span style="font-size:12px;"> : 让繁琐工作自动化 </span>

  </a>

      </h2>
      <div class="pub">
        
  
  [美] Al Sweigart / 王海鹏 / 人民邮电出版社 / 2021-3-1 / 89.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.2</span>

    <span class="pl">
        (91人评价)
    </span>
  </div>



        
  
  
  
    <p>本书是一本面向初学者的Python编程实用指南。本书不仅介绍了Python语言的基础知识，而且通过案例实践教读者如何使用这些知识和技能。本书的第一部分介绍了... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/35387685/buylinks">
        纸质版 89.00元
      </a>
    </span>

          </div>

            
            
  
    
    
  <div class="ebook-link">
    <a target="_blank" href="https://read.douban.com/ebook/336176935/?dcs=tag-buylink&amp;dcm=douban&amp;dct=35387685">去看电子版</a>
  </div>
  


      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/35635836/" 
  onclick="moreurl(this,{i:'18',query:'',subject_id:'35635836',from:'book_subject_search'})">
        <img class="" src="https://img1.doubanio.com/view/subject/s/public/s34071329.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/35635836/" title="Go语言设计与实现" 
  onclick="moreurl(this,{i:'18',query:'',subject_id:'35635836',from:'book_subject_search'})">

    Go语言设计与实现


    

  </a>

      </h2>
      <div class="pub">
        
  
  左书祺 / 人民邮电出版社 / 2021-11 / 139.80元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar40"></span>
        <span class="rating_nums">8.2</span>

    <span class="pl">
        (51人评价)
    </span>
  </div>



        
  
  
  
    <p>本书基于在读者之间广为传阅的同名开源电子书《Go语言设计与实现》，是难得一见的Go语言进阶图书。
书中结合近200幅生动的全彩图片，配上详尽的文字剖析与精选... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/35635836/buylinks">
        纸质版 117.10元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

      
      
  
  <li class="subject-item">
    <div class="pic">
      <a class="nbg" href="https://book.douban.com/subject/1052241/" 
  onclick="moreurl(this,{i:'19',query:'',subject_id:'1052241',from:'book_subject_search'})">
        <img class="" src="https://img2.doubanio.com/view/subject/s/public/s1074361.jpg"
          width="90">
      </a>
    </div>
    <div class="info">
      <h2 class="">
        
  
  <a href="https://book.douban.com/subject/1052241/" title="设计模式" 
  onclick="moreurl(this,{i:'19',query:'',subject_id:'1052241',from:'book_subject_search'})">

    设计模式


    
      <span style="font-size:12px;"> : 可复用面向对象软件的基础 </span>

  </a>

      </h2>
      <div class="pub">
        
  
  [美] Erich Gamma、Richard Helm、Ralph Johnson、John Vlissides / 李英军、马晓星、蔡敏、刘建中 等 / 机械工业出版社 / 2000-9 / 35.00元

      </div>


        
  
  
  
  <div class="star clearfix">
        <span class="allstar45"></span>
        <span class="rating_nums">9.1</span>

    <span class="pl">
        (3104人评价)
    </span>
  </div>



        
  
  
  
    <p>这本书结合设计实作例从面向对象的设计中精选出23个设计模式，总结了面向对象设计中最有价值的经验，并且用简洁可复用的形式表达出来。书中分类描述了一组设计良好、... </p>






      <div class="ft">
          
  <div class="collect-info">
  </div>


          <div class="cart-actions">
            
                
  

    <span class="buy-info">
      <a href="https://book.douban.com/subject/1052241/buylinks">
        纸质版 27.20元
      </a>
    </span>

          </div>

            
            
  

      </div>

    </div>
  </li>
  

  </ul>


      
    
    
    
        <div class="paginator">
        <span class="prev">
            &lt;前页
        </span>
        
        

                <span class="thispage">1</span>
                
            <a href="/tag/编程?start=20&amp;type=T" >2</a>
        
                
            <a href="/tag/编程?start=40&amp;type=T" >3</a>
        
                
            <a href="/tag/编程?start=60&amp;type=T" >4</a>
        
                
            <a href="/tag/编程?start=80&amp;type=T" >5</a>
        
                
            <a href="/tag/编程?start=100&amp;type=T" >6</a>
        
                
            <a href="/tag/编程?start=120&amp;type=T" >7</a>
        
                
            <a href="/tag/编程?start=140&amp;type=T" >8</a>
        
                
            <a href="/tag/编程?start=160&amp;type=T" >9</a>
        
            <span class="break">...</span>
                
            <a href="/tag/编程?start=1080&amp;type=T" >55</a>
        
            <a href="/tag/编程?start=1100&amp;type=T" >56</a>
        
        <span class="next">
            <link rel="next" href="/tag/编程?start=20&amp;type=T"/>
            <a href="/tag/编程?start=20&amp;type=T" >后页&gt;</a>
        </span>

        </div>


  </div>
</div>
      <div class="aside">
        
  <br/>
  <div id="dale_book_tag_top_right"></div>

    
  

  <h2>
    <span>相关的标签</span>
      &nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;

  </h2>


    
  <div class="tags-list">
      <a href="/tag/Programming">Programming</a>
      <a href="/tag/计算机">计算机</a>
      <a href="/tag/程序设计">程序设计</a>
      <a href="/tag/C++">C++</a>
      <a href="/tag/programming">programming</a>
      <a href="/tag/软件开发">软件开发</a>
      <a href="/tag/计算机科学">计算机科学</a>
      <a href="/tag/C#">C#</a>
  </div>


  
  <form name="tsp_form" action="/tag/" method="GET">
    <input name="search_text" class="j a_search_text" type="text" size="24" maxlength="36" title="去其他标签" value=""/>
    <input class="butt" type="submit" value="进入"/>
  </form>
  <br/>
  <br/>


  <p class="pl2 mb30">
    &gt; <a href="/tag/">浏览全部图书标签</a>
  </p>

    
  

  <h2>
    <span>&#34;编程&#34;相关豆列</span>
      &nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;&nbsp;&middot;

  </h2>


    

<ul class="bs">
  <li>
    <a href="https://www.douban.com/doulist/336424/">编程C语言,Python等类书籍+知乎话题</a>
      <span class="pl">(IVAN)</span>
  </li>
  <li>
    <a href="https://www.douban.com/doulist/430092/">数学计算机专业书籍</a>
      <span class="pl">(万籁君)</span>
  </li>
  <li>
    <a href="https://www.douban.com/doulist/172288/">计算机理论</a>
      <span class="pl">(mashan_snail)</span>
  </li>
  <li>
    <a href="https://www.douban.com/doulist/85050/">【C/C++学习指南】</a>
      <span class="pl">(小李飞刀)</span>
  </li>
  <li>
    <a href="https://www.douban.com/doulist/1525602/">Hack的技术与艺术</a>
      <span class="pl">(lyb)</span>
  </li>
</ul>


    <br/>
    <br/>
    
<div class="block5 movie_show">
    <h2>最近受关注的书-编程</h2>
    <div class="content clearfix" id="book_rec">
            <dl>
                <dt><a onclick="moreurl(this, {'r1001':'book-tag-top'})" 
                href="https://book.douban.com/subject/25815142/"><img src="https://img3.doubanio.com/view/subject/s/public/s27215120.jpg" 
                class="m_sub_img" /></a></dt>
                <dd><a onclick="moreurl(this, {'r1001':'book-tag-top'})"
                href="https://book.douban.com/subject/25815142/">游戏引擎架构</a>
                </dd>
            </dl>
            <dl>
                <dt><a onclick="moreurl(this, {'r1001':'book-tag-top'})" 
                href="https://book.douban.com/subject/11614538/"><img src="https://img2.doubanio.com/view/subject/s/public/s11194203.jpg" 
                class="m_sub_img" /></a></dt>
                <dd><a onclick="moreurl(this, {'r1001':'book-tag-top'})"
                href="https://book.douban.com/subject/11614538/">程序员的职业素养</a>
                </dd>
            </dl>
            <dl>
                <dt><a onclick="moreurl(this, {'r1001':'book-tag-top'})" 
                href="https://book.douban.com/subject/27082348/"><img src="https://img9.doubanio.com/view/subject/s/public/s29488325.jpg" 
                class="m_sub_img" /></a></dt>
                <dd><a onclick="moreurl(this, {'r1001':'book-tag-top'})"
                href="https://book.douban.com/subject/27082348/">自己动手写Docker</a>
                </dd>
            </dl>
            <div class="clearfix rr" style="width:100%"></div>
            <dl>
                <dt><a onclick="moreurl(this, {'r1001':'book-tag-top'})" 
                href="https://book.douban.com/subject/1782316/"><img src="https://img2.doubanio.com/view/subject/s/public/s4567393.jpg" 
                class="m_sub_img" /></a></dt>
                <dd><a onclick="moreurl(this, {'r1001':'book-tag-top'})"
                href="https://book.douban.com/subject/1782316/">Concepts, Techniques, and Models of Computer Programming</a>
                </dd>
            </dl>
            <dl>
                <dt><a onclick="moreurl(this, {'r1001':'book-tag-top'})" 
                href="https://book.douban.com/subject/26616762/"><img src="https://img1.doubanio.com/view/subject/s/public/s29792058.jpg" 
                class="m_sub_img" /></a></dt>
                <dd><a onclick="moreurl(this, {'r1001':'book-tag-top'})"
                href="https://book.douban.com/subject/26616762/">Programming Rust</a>
                </dd>
            </dl>
            <dl>
                <dt><a onclick="moreurl(this, {'r1001':'book-tag-top'})" 
                href="https://book.douban.com/subject/35720728/"><img src="https://img3.doubanio.com/view/subject/s/public/s34082690.jpg" 
                class="m_sub_img" /></a></dt>
                <dd><a onclick="moreurl(this, {'r1001':'book-tag-top'})"
                href="https://book.douban.com/subject/35720728/">Go语言精进之路</a>
                </dd>
            </dl>
            <div class="clearfix rr" style="width:100%"></div>
    </div>
</div>



  




    
<script type="text/javascript">
    (function (global) {
        var newNode = global.document.createElement('script'),
            existingNode = global.document.getElementsByTagName('script')[0],
            adSource = '//erebor.douban.com/',
            userId = '',
            browserId = 'UbmE-XEY7OI',
            criteria = '3:/tag/编程?start=0&amp;type=T',
            preview = '',
            debug = false,
            adSlots = ['dale_book_tag_top_right'];

        global.DoubanAdRequest = {src: adSource, uid: userId, bid: browserId, crtr: criteria, prv: preview, debug: debug};
        global.DoubanAdSlots = (global.DoubanAdSlots || []).concat(adSlots);

        newNode.setAttribute('type', 'text/javascript');
        newNode.setAttribute('src', '//img1.doubanio.com/MnJwNGlleS9mL2FkanMvY2M1OGQyNTQ2N2I2YmQzOTlmNTliMGJiMjI4MWRhZTlkNTNjYTFkZC9hZC5yZWxlYXNlLmpz');
        newNode.setAttribute('async', true);
        existingNode.parentNode.insertBefore(newNode, existingNode);
    })(this);
</script>











      </div>
      <div class="extra">
        
      </div>
    </div>
  </div>

        
  <div id="footer">
    
<span id="icp" class="fleft gray-link">
    &copy; 2005－2022 douban.com, all rights reserved 北京豆网科技有限公司
</span>

<a href="https://www.douban.com/hnypt/variformcyst.py" style="display: none;"></a>

<span class="fright">
    <a href="https://www.douban.com/about">关于豆瓣</a>
    · <a href="https://www.douban.com/jobs">在豆瓣工作</a>
    · <a href="https://www.douban.com/about?topic=contactus">联系我们</a>
    · <a href="https://www.douban.com/about/legal">法律声明</a>
    
    · <a href="https://help.douban.com/?app=book" target="_blank">帮助中心</a>
    · <a href="https://book.douban.com/library_invitation">图书馆合作</a>
    · <a href="https://www.douban.com/doubanapp/">移动应用</a>
    · <a href="https://www.douban.com/partner/">豆瓣广告</a>
</span>

  </div>

    </div>
      

    <!-- COLLECTED JS -->
    <!-- mako -->
    
  <script>
    $(function () {
      function update_tags(current) {
        html = $.map(current, function(o) {return '<span class="tags-name">' + o + '</span><sup class="tags-del"></sup> '}).join('<span class="tags-add">+</span>');
        $('h1').html('<span class="name">电影标签：</span>' + html);
      };

      function get_current_tags() {
        var r = [];
        $('.tags-name').each(function() {r.push($(this).text())});
        return r;
      };

      function get_current_sort_type() {
        url = location.href;
        array = url.split("?")
        if (array.length == 1) {
          return "";
        }
        query = array[1];
        arr = query.split("&");
        for (var i = 0; i < arr.length; i++) {
          arr2 = arr[i].split('=');
          if (arr2[0] == "type") {
            if (arr2.length == 1) {
              return "";
            }
            return arr2[1];
          }
        }
        return "";
      };

      function update_subject_list(current, sort_type) {
        $('#subject_list').html('<span class="pl">载入中，请稍候...</span>');
        $('#subject_list').load_withck('/j/tag/j_subject_list', {tags:current.join(' '), type:sort_type}, function() {});
      };

      function update_related_tags_and_subject(current, sort_type) {
        update_related_tags(current, sort_type);
        // update_subject_list(current, sort_type);
      };

      function update_related_tags(current, sort_type) {
        $('#related_tags').load_withck('/j/tag/j_related_tag', {tags:current.join(' '), type:sort_type}, function() {
          $('.more').toggle(
            function () {
              $('.tags-hide').css('display', 'inline');
              $(this).text('收起▲');
            },
            function () {
              $('.tags-hide').hide();
              $(this).text('更多▼');
            }
          );
          $('.tags-del').click(function () {
            if($(this).prev().prev().text() == '+') {
              $(this).prev().prev().remove();
            } else if ($(this).next().text() == '+') {
              $(this).next().remove();
            }
            $(this).prev().remove().end().remove();
            current = get_current_tags();
            if (current.length == 0){
              document.location.href = '/tag/';
            } else {
              //setHash(current.join(' '));
              //location.href = '/tag/' + (location.href).split('#')[1] + '?type=' + sort_type;
              sort_type = get_current_sort_type();
              location.href = '/tag/' + current.join(' ') + '?type=' + sort_type;
            }
          });
          $('.add-tag').click(function () {
            current = get_current_tags();
            current.push($(this).text());
            update_tags(current);
            //setHash(current.join(' '));
            //location.href = '/tag/' + (location.href).split('#')[1] + '?type=' + sort_type;
            sort_type = get_current_sort_type();
            location.href = '/tag/' + current.join(' ') + '?type=' + sort_type;
          });
          $('.tagsInput').keypress(function (e) {
            if (e.which == 13) {
              current = get_current_tags();
              current.push($(this).val());
              update_tags(current);
              sort_type = get_current_sort_type();
              //setHash(current.join(' '));
              //location.href = '/tag/' + (location.href).split('#')[1] + '?type=' + sort_type;
              location.href = '/tag/' + current.join(' ') + '?type=' + sort_type;
            }
          });
          checkInput($('.tagsInput'));
        });
      };

      current = get_current_tags();
      sort_type = get_current_sort_type();
      update_related_tags_and_subject(current, sort_type);

      function checkInput (o) {
        if (!o.val() || o.val() == o.attr('title')) {
          o.addClass('greyinput');
          o.val(o.attr('title'));
        }
        o.focus(function(){
          o.removeClass('greyinput');
          if(o.val() == o.attr('title')) o.val('');
        });
        o.blur(function(){
          if(!o.val()){
            o.addClass('greyinput');
            o.val(o.attr('title'));
          }
        });
      }

      if($.browser.msie && $.browser.version == 6.0) {
        $('.tags-del').hover(
          function () { $(this).addClass('tags-hover'); },
          function () { $(this).removeClass('tags-hover'); }
        );
      }
      if (location.href.split('#').length > 1) {
        //changeUrl ();
      }

  });

  function setHash (a) {
    $.browser.msie?$.locationHash(a):location.hash = a;
  }
  </script>

    
  

<script type="text/javascript">
  var _paq = _paq || [];
  _paq.push(['trackPageView']);
  _paq.push(['enableLinkTracking']);
  (function() {
    var p=(('https:' == document.location.protocol) ? 'https' : 'http'), u=p+'://fundin.douban.com/';
    _paq.push(['setTrackerUrl', u+'piwik']);
    _paq.push(['setSiteId', '100001']);
    var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0]; 
    g.type='text/javascript';
    g.defer=true; 
    g.async=true; 
    g.src=p+'://s.doubanio.com/dae/fundin/piwik.js';
    s.parentNode.insertBefore(g,s);
  })();
</script>

<script type="text/javascript">
var setMethodWithNs = function(namespace) {
  var ns = namespace ? namespace + '.' : ''
    , fn = function(string) {
        if(!ns) {return string}
        return ns + string
      }
  return fn
}

var gaWithNamespace = function(fn, namespace) {
  var method = setMethodWithNs(namespace)
  fn.call(this, method)
}

var _gaq = _gaq || []
  , accounts = [
      { id: 'UA-7019765-1', namespace: 'douban' }
    , { id: 'UA-7019765-16', namespace: '' }
    ]
  , gaInit = function(account) {
      gaWithNamespace(function(method) {
        gaInitFn.call(this, method, account)
      }, account.namespace)
    }
  , gaInitFn = function(method, account) {
      _gaq.push([method('_setAccount'), account.id])

      
  _gaq.push([method('_addOrganic'), 'google', 'q'])
  _gaq.push([method('_addOrganic'), 'baidu', 'wd'])
  _gaq.push([method('_addOrganic'), 'soso', 'w'])
  _gaq.push([method('_addOrganic'), 'youdao', 'q'])
  _gaq.push([method('_addOrganic'), 'so.360.cn', 'q'])
  _gaq.push([method('_addOrganic'), 'sogou', 'query'])
  if (account.namespace) {
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣'])
    _gaq.push([method('_addIgnoredOrganic'), 'douban'])
    _gaq.push([method('_addIgnoredOrganic'), '豆瓣网'])
    _gaq.push([method('_addIgnoredOrganic'), 'www.douban.com'])
  }

      if (account.namespace === 'douban') {
        _gaq.push([method('_setDomainName'), '.douban.com'])
      }

        _gaq.push([method('_setCustomVar'), 1, 'responsive_view_mode', 'desktop', 3])

        _gaq.push([method('_setCustomVar'), 2, 'login_status', '0', 2]);

      _gaq.push([method('_trackPageview')])
    }

for(var i = 0, l = accounts.length; i < l; i++) {
  var account = accounts[i]
  gaInit(account)
}


;(function() {
    var ga = document.createElement('script');
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
    ga.setAttribute('async', 'true');
    document.documentElement.firstChild.appendChild(ga);
})()
</script>

    <!-- dae-web-book--default-5d9ddc57cd-9swtx-->

</body>
</html>




































`
	pages := models.GetTotalPages(bookHtml)

	bookurl := models.GetBookUrl(bookHtml)

	c.Ctx.WriteString(fmt.Sprintf("当前共 %d 页\n 每页共有 %d 本书",pages,len(bookurl)))
	for _,v := range bookurl {
		c.Ctx.WriteString(fmt.Sprintf("书的URL %s\n ",v))
	}

}
