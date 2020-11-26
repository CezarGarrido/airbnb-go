function NewSharedView() {
  
    function showNavbar(element) {
      element.innerHTML = navbar();
    }

    function showFooter(element) {
      element.innerHTML = footer();
    }


    return {
        showNavbar: showNavbar,
        showFooter: showFooter,
    };
  }
  
  const navbar = () => `
  <header class="simple-header for-sticky white">
  <div class="menu">
      <div class="container">
          <div class="logo">
              <a href="index.html" title="">
                  <i class="fa fa-get-pocket"></i>
                  <span>Airbnb</span>
                  <strong></strong>
              </a>
          </div><!-- LOGO -->
          <div class="popup-client">
              <span><i class="fa fa-user"></i> / Acessar Conta</span>
          </div>
          <span class="menu-toggle"><i class="fa fa-bars"></i></span>
          <nav>
              <h1 class="nocontent outline">--- Main Navigation ---</h1>

              <ul>
                  <li class="menu-item-has-children">
                      <a href="#index.html" title="">HOME</a>
                  </li>
                  <li class="menu-item-has-children">
                      <a href="#vehiculs3.html" title="">Locations</a>
                  </li>
              </ul>
          </nav>

      </div>
  </div>
</header>
  `;
  
  const footer = () => `
  <footer class="light-footer">
  <section class="top-line">
      <div style="background: url(img/footer.jpg) repeat scroll 50% 422.28px transparent;"
          class="parallax scrolly-invisible no-parallax whitish"></div><!-- PARALLAX BACKGROUND IMAGE -->
      <div class="container">
          <div class="row">
              <div class="col-md-3 column">
                  <div class="links_widget widget">
                      <div class="heading1">
                          <h2><span>Useful</span> links</h2>
                      </div><!-- heading -->
                      <ul>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i> Home</a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i> About us</a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i> Services</a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i> Services</a></li>
                      </ul>
                  </div>
              </div>

              <div class="col-md-3 column">
                  <div class="links_widget widget">
                      <div class="heading1">
                          <h2><span>Useful</span> Places</h2>
                      </div><!-- heading -->
                      <ul>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i> Lorem ipsum </a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i> Closest ipsum </a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i> Lorem dolom </a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i> Positioned ipl</a></li>
                      </ul>
                  </div>
              </div>

              <div class="col-md-3 column">
                  <div class="subscribe_widget widget">
                      <div class="heading1">
                          <h2><span>Subscribe</span> Us</h2>
                      </div><!-- heading -->
                      <p>Positioning the closest positioned for abs positioning</p>
                      <form>
                          <label><input type="text" placeholder="TYPE YOUR EMAIL" /></label>
                          <button type="submit" class="flat-btn"><i class="ti ti-email"></i></button>
                      </form>
                  </div>
              </div>

          </div>
      </div>
  </section>
  <div class="bottom-line">
      <div class="container">
          <span>Copyright All Right Reserved 2016 <a href="#" title="">KimaroTec</a></span>
          <ul>
              <li><a title="" href="#">HOME</a></li>
              <li><a title="" href="#">ABOUT</a></li>
              <li><a title="" href="#">VEHICULS</a></li>
              <li><a title="" href="#">BLOG</a></li>
              <li><a title="" href="#">CONTACT</a></li>
          </ul>
      </div>
  </div>
  <a href="#" class="scrollToTop"><i class="ti ti-arrow-circle-up"></i></a>
</footer>
  `;
  