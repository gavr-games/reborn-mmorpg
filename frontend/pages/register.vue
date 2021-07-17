<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <div class="container mt-20 small">
      <div class="inner rpgui-container framed">
        <header>
          <h1 class="main-title">REBORN</h1>
		      <hr class="golden" />
          <h2>Register</h2>
        </header>

        <div class="rpgui-center">
            <label>Your hero name:</label>
            <input type="text" v-model="username" placeholder="myusername">
            <br/><br/>

            <label>Your email:</label>
            <input type="text" v-model="email" placeholder="myemail@example.com">
            <br /><br />

            <label>Your password:</label>
            <input type="password" v-model="password">
            <br/><br/>
		    </div>
        <br/><br/>
        <div class="rpgui-center">
          <a @click="register()"><button type="button" class="rpgui-button golden"><p>Register</p></button></a>
          <br />
          <NuxtLink to="/"><button type="button" class="rpgui-button"><p>Back</p></button></NuxtLink>
        </div>
        <br /><br />
			  <p v-show="showSuccessMessage">
          Thank you for registration, you can now <NuxtLink to="/login">login</NuxtLink>.
			  </p>
        <p v-show="showErrorMessage" class="error">
          {{ errorMessage }}
			  </p>
		    <br /><br />
        <hr style="clear:both">
		    <br /><br />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      email: '',
      password: '',
      errorMessage: '',
      showSuccessMessage: false,
      showErrorMessage: false,
    }
  },
  methods: {
    async register() {
      this.showSuccessMessage = false
      this.showErrorMessage = false
      const context = this
      this.$axios.$post('/players', {
        username: this.username,
        email: this.email,
        password: this.password
      })
      .then(response => {
        context.showSuccessMessage = true
        context.username = ''
        context.email = ''
        context.password = ''
        console.log(response)
      })
      .catch(error => {
        if (error.response) {
          context.errorMessage = error.response.data.error
          context.showErrorMessage = true
        }
      });
    }
  }
}
</script>
