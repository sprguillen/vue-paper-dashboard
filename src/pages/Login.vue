<template>
  <section class="login-section">
    <div class="left">
      <div class="logo-container">
        <img class="login-logo" src="../assets/img/logo.png" /><br/><br/><br/>
        <p>Automatically personalize your website & email to each visitor</p>
      </div>
    </div>
    <div class="right">
      <header class="sign-up-header">
        <span>Don't have an account? <a href="#">Signup</a></span>
      </header>
      <section class="sign-up-form-container">
        <header>
          <img class="plants-left d-none d-md-block" src="../assets/img/plants-left.svg" />
          <span>Sign in to DataCue</span>
          <hr />
        </header>
        <section>
          <form @submit.prevent="onSubmit">
            <div class="form-group">
              <label for="email-input">Email address</label>
              <input id="email-input" type="email" class="form-control bg-grey" v-model="form.email">
            </div>
            <div class="form-group">
              <label for="pass-input">Password</label>
              <label class="show-password d-none d-md-block" @click="changePassFieldType" for="pass-input">
                <i class="ti-eye"/> Show password
              </label>
              <input id="pass-input" :type="passFieldType" class="form-control bg-grey" v-model="form.pass">
              <label class="show-password d-md-none" @click="changePassFieldType" for="pass-input">
                <i class="ti-eye"/> Show password
              </label>
            </div>
            <div>
              <button type="submit" class="btn btn-custom" :disabled="disableSubmit">Submit</button>
              <img class="plants-right d-none d-md-block" src="../assets/img/plants-right.svg" />
            </div>
          </form>
        </section>
      </section>
    </div>
  </section>
</template>
<script>
export default {
  name: 'Login',
  data() {
    return {
      form: {
        email: '',
        pass: ''
      },
      passFieldType: 'password'
    }
  },
  computed: {
    disableSubmit() {
      let hasEmail = this.form.email.length > 0;
      let hasPass = this.form.pass.length > 0;

      return !(hasEmail && hasPass);
    }
  },
  methods: {
    changePassFieldType() {
      if (this.passFieldType === 'password') {
        this.passFieldType = 'text';
      } else {
        this.passFieldType = 'password';
      }
    },
    onSubmit() {
      console.log('here');
    }
  }
}
</script>
<style>
  .sign-up-header > span {
    font-size: 1rem;
  }

  .sign-up-header > span > a {
    color: #ffba00;
  }

  .show-password {
    float: right;
    cursor: pointer;
  }

  .btn-custom {
    background-color: #ffba00;
    border: none;
  }

  .bg-grey {
    background-color: #F0F0F0;
  }

  @media (min-width: 768px) and (max-width: 779px) {
    .show-password {
      float: left;
    }
  }

  @media (min-width: 768px) {
    .login-section {
      height: 100vh;
      margin: 0 auto;
      width: 60%;
    }

    .left {
      float: left;
      width: 50%;
      height: 100%;
      background: url('../assets/img/squares-background.svg');
      background-size: 100% auto;
      display: flex;
      justify-content: center;
      align-items: center;
      text-align: center;
      display: table;
    }

    .logo-container {
      display: table-cell;
      vertical-align: middle;
    }

    .right {
      float: left;
      position: relative;
      width: 50%;
    }

    .logo-container > p {
      font-size: 1.5rem;
      width: 70%;
      margin: 0 auto;
    }

    .login-logo {
      width: 50%;
      height: 9%;
    }

    .sign-up-header {
      position: absolute;
      right: 0;
      top: 1.5rem;
    }

    .sign-up-form-container {
      width: 80%;
      position: absolute;
      top: 4rem;
      left: 20%;
    }

    .sign-up-form-container > header {
      text-align: center;
    }

    .plants-left {
      width: 6rem;
      height: 6rem;
      float: left;
    }

    .plants-right {
      width: 6rem;
      height: 6rem;
      float: right;
    }

    .sign-up-form-container > header > span {
      font-size: 1.2rem;
      position: relative;
      bottom: -4rem;
      font-weight: 600;
    }

    .sign-up-form-container > header > hr {
      position: relative;
      top: 4rem;
    }

    .sign-up-form-container > section {
      margin-top: 5rem;
    }

    .btn-custom {
      position: absolute;
      bottom: 0;
    }
  }

  @media (max-width: 767px) {
    .login-section {
      height: 100vh;
      margin: 0 auto;
      width: 60%;
    }

    .left {
      width: 100%;
      height: 400px;
      background: url('../assets/img/background-with-squares.svg');
      background-size: 100% auto;
      display: flex;
      justify-content: center;
      align-items: center;
      text-align: center;
      display: table;
      position: relative;
    }

    .right {
      position: relative;
      top: 2rem;
    }

    .logo-container {
      position: absolute;
      bottom: 1rem;
    }

    .logo-container > p {
      font-size: 1.3rem;
      width: 70%;
      margin: 0 auto;
    }

    .login-logo {
      width: 50%;
      height: 9%;
    }

    .sign-up-header {
      text-align: center;
    }

    .sign-up-form-container {
      margin-top: 3rem;
    }

    .sign-up-form-container > header {
      text-align: center;
    }

    .sign-up-form-container > header > span {
      font-size: 1.2rem;
      font-weight: 600;
    }

    .sign-up-form-container > header > hr {
      width: 40%;
    }
  }
</style>
