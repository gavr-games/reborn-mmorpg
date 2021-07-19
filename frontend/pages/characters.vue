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

        <div class="row">
          <div class="col-3 rpgui-container framed-golden-2" v-for="char in characters" :key="char.id" @click="selectCharacter(char.id)">
            <div class="rpgui-icon helmet-slot add rpgui-cursor-point"></div>
            {{ char.name }}
          </div>
          <div class="col-3 rpgui-container framed-golden-2" v-for="index in (4 - characters.length)" :key="'index'+index">
            <div class="rpgui-icon empty-slot add rpgui-cursor-point" @click="displayCreateForm"><p>+</p></div>
          </div>
        </div>

        <div class="rpgui-center" v-if="showCreateForm">
          <div class="add-character-form">
            <label>Your character name:</label>
            <input type="text" v-model="name" placeholder="myhero">
            <br/><br/>

            <label>Choose appearance:</label>
            <select class="rpgui-dropdown" data-rpguitype="dropdown" v-model="gender">
              <option value="male">Male</option>
              <option value="female">Female</option>
            </select>
            <br /><br />
            <button type="button" class="rpgui-button golden" @click="create"><p>Create</p></button>
          </div>
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
      characters: [],
      name: '',
      gender: 'male',
      errorMessage: '',
      showErrorMessage: false,
      showCreateForm: false,
    }
  },

  mounted() {
    if (!this.$auth.loggedIn) {
      this.$router.push('login')
    } else {
      this.getList()
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
  },

  methods: {
    displayCreateForm() {
      this.showCreateForm = true;
    },
    create() {
      this.showErrorMessage = false
      const context = this
      this.$axios.$post('/characters', {
        name: this.name,
        gender: this.gender
      })
      .then(response => {
        context.name = ''
        context.gender = 'male'
        context.showCreateForm = false
        context.characters.push(response)
      })
      .catch(error => {
        if (error.response) {
          context.errorMessage = error.response.data.error
          context.showErrorMessage = true
        }
      });
    },
    getList() {
      this.$axios.$get('/characters')
      .then(response => {
        this.characters = response
      })
    },
    selectCharacter(id) {
      this.$store.commit('characters/set', id)
      this.$router.push('/game');
    }
  }
}
</script>
