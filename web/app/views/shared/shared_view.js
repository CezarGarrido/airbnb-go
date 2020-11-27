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
      <div style="background: url(&quot;img/footer.jpg;) 50% 0px repeat scroll transparent;" class="parallax scrolly-invisible no-parallax whitish"></div><!-- PARALLAX BACKGROUND IMAGE -->
      <div class="container">
          <div class="row">
              <div class="col-md-3 column">
                  <div class="links_widget widget">
                      <div class="heading1">
                          <h2><span><font style="vertical-align: inherit;"></font></span><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Links </font><span><font style="vertical-align: inherit;">úteis</font></span></font></h2>
                      </div><!-- heading -->
                      <ul>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> Casa</font></font></a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> Sobre nós</font></font></a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> Serviços</font></font></a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> Serviços</font></font></a></li>
                      </ul>
                  </div>
              </div>

              <div class="col-md-3 column">
                  <div class="links_widget widget">
                      <div class="heading1">
                          <h2><span><font style="vertical-align: inherit;"></font></span><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Lugares </font><span><font style="vertical-align: inherit;">Úteis</font></span></font></h2>
                      </div><!-- heading -->
                      <ul>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> Lorem ipsum </font></font></a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> Ipsum mais próximo </font></font></a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> Lorem Dolom </font></font></a></li>
                          <li><a href="#" title=""><i class="fa fa-angle-right"></i><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> Ipl posicionado</font></font></a></li>
                      </ul>
                  </div>
              </div>

              <div class="col-md-3 column">
                  <div class="subscribe_widget widget">
                      <div class="heading1">
                          <h2><span><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Inscreva-</font></font></span><font style="vertical-align: inherit;"><font style="vertical-align: inherit;"> se</font></font></h2>
                      </div><!-- heading -->
                      <p><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Posicionando o mais próximo posicionado para posicionamento de abs</font></font></p>
                      <form>
                          <label><input type="text" placeholder="DIGITE SEU E-MAIL"></label>
                          <button type="submit" class="flat-btn"><i class="ti ti-email"></i></button>
                      </form>
                  </div>
              </div>

          </div>
      </div>
  </section>
  <div class="bottom-line">
      <div class="container">
          <span><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Copyright Todos os direitos reservados 2020</font></font></span>
          <ul>
              <li><a title="" href="#"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">INICIO</font></font></a></li>
              <li><a title="" href="#"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">SOBRE</font></font></a></li>
              <li><a title="" href="#"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">APARTAMENTOS</font></font></a></li>
              <li><a title="" href="#"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">CONTATO</font></font></a></li>
          </ul>
      </div>
  </div>
  <a href="#" class="scrollToTop" style="display: inline;"><i class="ti ti-arrow-circle-up"></i></a>
</footer>
  `;
  