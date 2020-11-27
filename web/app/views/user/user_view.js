function NewUserView() {

    function AccountPopupComponent(element) {
        element.innerHTML = accountPopupComponent();
    }

    return {
        AccountPopupComponent: AccountPopupComponent,
    };
}
const accountPopupComponent = () => `

  <div class="account-popup-sec">
  <div class="account-popup-area">
      <div class="account-popup">
          <div class="row">
              <div class="col-md-6">
                  <div class="account-user">
                      <div class="logo">
                          <a href="#" title="">
                              <i class="fa fa-get-pocket"></i>
                              <span>Airbnb</span>
                              <strong></strong>
                          </a>
                      </div><!-- LOGO -->
                      <form>
                          <h1>Login</h1>
                          <div class="field">
                              <input type="text" id="input-email" placeholder="Email" />
                          </div>
                          <div class="field">
                              <input type="password" placeholder="Senha" />
                          </div>
                          <div class="field">
                              <input type="submit" value="Enviar" class="flat-btn" />
                          </div>
                      </form>
                      <i></i>
                      <span></span>
                      <ul class="social-btns">
                          <li></li>
                          <li></li>
                          <li></li>
                      </ul>
                  </div>
              </div>
              <div class="col-md-6">
                  <div class="registration-sec">
                      <h1>Criar Conta</h1>
                      <form>
                          <div class="field">
                              <input type="text" id="input-register-name" placeholder="Nome" />
                          </div>
                          <div class="field">
                              <input type="text" id="input-register-email" placeholder="Email" />
                          </div>
                          <div class="field">
                              <input type="password" id="input-register-password" placeholder="Senha" />
                          </div>
                          <div class="field">
                              <input type="password" id="input-register-repeat-password" placeholder="Repetir Senha" />
                          </div>
                          <input type="button" id="btn-resgister-user" value="Criar Agora" class="flat-btn" />
                      </form>
                  </div><!-- Registration sec -->
              </div>
          </div>
          <span class="close-popup"><i class="fa fa-close"></i></span>
      </div>
  </div>
</div><!-- Account Popup Sec -->

`;