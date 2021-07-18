<template>
  <div class="rpgui-content" style="overflow-y:scroll;">
    <div class="container mt-20">
      <div class="inner rpgui-container framed">
        <header>
          <h1 class="main-title">REBORN</h1>
		      <hr class="golden" />
          <h2>Hello {{ username }}, please select the character</h2>
        </header>

        <div class="rpgui-center">
          <NuxtLink to="/"><button type="button" class="rpgui-button"><p>Home</p></button></NuxtLink>
		    </div>
		    <br /><br />

		    <br /><br /><br /><br />
        <hr style="clear:both">
		    <br /><br />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  mounted() {
    if (!this.$auth.loggedIn) {
      this.$router.push('login')
    }
  },

  computed: {
    username() {
      if (this.$auth.user) {
        return this.$auth.user.username
      } else {
        ''
      }
    }
  }
}
</script>
