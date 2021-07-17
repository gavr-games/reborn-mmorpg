<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <div class="container mt-20 small">
      <div class="inner rpgui-container framed">
        <header>
          <h1 class="main-title">REBORN</h1>
		      <hr class="golden" />
          <h2>Login</h2>
        </header>

        <div class="rpgui-center">
            <label>Your username:</label>
            <input type="text" v-model="username" placeholder="myusername">
            <br/><br/>

            <label>Your password:</label>
            <input type="password" v-model="password">
            <br/><br/>
		    </div>
        <br/><br/>
        <div class="rpgui-center">
          <a @click="login()"><button type="button" class="rpgui-button golden"><p>Login</p></button></a>
          <br />
          <NuxtLink to="/"><button type="button" class="rpgui-button"><p>Back</p></button></NuxtLink>
        </div>
        <br /><br />
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
      password: '',
      errorMessage: '',
      showErrorMessage: false,
    }
  },
  methods: {
    async login() {
      this.showErrorMessage = false
      const context = this
      this.$axios.$post('/login', {
        username: this.username,
        password: this.password
      })
      .then(response => {

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
