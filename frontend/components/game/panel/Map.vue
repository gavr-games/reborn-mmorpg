<template>
  <div :class="`game-panel map ${showMapPanel ? '' : 'hide-map'}`">
    <div class="game-panel-content">
      <div class="pin" :style="{ left: left + 'px', top: top + 'px' }"></div>
      <img src="/engine_api/maps/floor_0_map.jpg" alt="Floor map" id="floor-map" /> <br />
      <button type="button" class="rpgui-button" @click="showMapPanel = false"><p>Close</p></button>
    </div>
  </div>
</template>

<script>
import { EventBus } from "~/plugins/game/event_bus";

export default {
  data() {
    return {
      showMapPanel: false,
      left: 0,
      top: 0,
    }
  },

  created() {
    EventBus.$on("show-map", this.showMap)
    EventBus.$on("update_object", this.updateObject)
    EventBus.$on("add_object", this.updateObject)
  },

  beforeDestroy() {
    EventBus.$off("show-map", this.showMap)
    EventBus.$off("update_object", this.updateObject)
    EventBus.$off("add_object", this.updateObject)
  },

  methods: {
    showMap() {
      this.showMapPanel = true
    },
    updateObject(obj) {
      if (obj.Properties["player_id"] === this.$store.state.characters.selectedCharacterId) {
        if (document.getElementById("floor-map")) {
          this.left = obj.X -3
          this.top = document.getElementById("floor-map").height - obj.Y -24
        }
      }
    },
  }
}
</script>


<style>
.hide-map {
  display:none;
}
.map {
  position: absolute;
  top: 100px;
}
.pin {
  width: 30px;
  height: 30px;
  opacity: 0.7;
  border-radius: 50% 50% 50% 0;
  background: #89849b;
  position: absolute;
  -webkit-transform: rotate(-45deg);
  -moz-transform: rotate(-45deg);
  -o-transform: rotate(-45deg);
  -ms-transform: rotate(-45deg);
  transform: rotate(-45deg);
  left: 0;
  top: 0;
  -webkit-animation-name: bounce;
  -moz-animation-name: bounce;
  -o-animation-name: bounce;
  -ms-animation-name: bounce;
  animation-name: bounce;
  -webkit-animation-fill-mode: both;
  -moz-animation-fill-mode: both;
  -o-animation-fill-mode: both;
  -ms-animation-fill-mode: both;
  animation-fill-mode: both;
  -webkit-animation-duration: 1s;
  -moz-animation-duration: 1s;
  -o-animation-duration: 1s;
  -ms-animation-duration: 1s;
  animation-duration: 1s;
}
.pin:after {
  content: '';
  width: 14px;
  height: 14px;
  margin: 8px 0 0 8px;
  background: #2f2f2f;
  position: absolute;
  border-radius: 50%;
}

@-moz-keyframes bounce {
  0% {
    opacity: 0;
    -ms-filter: "progid:DXImageTransform.Microsoft.Alpha(Opacity=0)";
    filter: alpha(opacity=0);
    -webkit-transform: translateY(-2000px) rotate(-45deg);
    -moz-transform: translateY(-2000px) rotate(-45deg);
    -o-transform: translateY(-2000px) rotate(-45deg);
    -ms-transform: translateY(-2000px) rotate(-45deg);
    transform: translateY(-2000px) rotate(-45deg);
  }
  60% {
    opacity: 1;
    -ms-filter: none;
    filter: none;
    -webkit-transform: translateY(30px) rotate(-45deg);
    -moz-transform: translateY(30px) rotate(-45deg);
    -o-transform: translateY(30px) rotate(-45deg);
    -ms-transform: translateY(30px) rotate(-45deg);
    transform: translateY(30px) rotate(-45deg);
  }
  80% {
    -webkit-transform: translateY(-10px) rotate(-45deg);
    -moz-transform: translateY(-10px) rotate(-45deg);
    -o-transform: translateY(-10px) rotate(-45deg);
    -ms-transform: translateY(-10px) rotate(-45deg);
    transform: translateY(-10px) rotate(-45deg);
  }
  100% {
    -webkit-transform: translateY(0) rotate(-45deg);
    -moz-transform: translateY(0) rotate(-45deg);
    -o-transform: translateY(0) rotate(-45deg);
    -ms-transform: translateY(0) rotate(-45deg);
    transform: translateY(0) rotate(-45deg);
  }
}
@-webkit-keyframes bounce {
  0% {
    opacity: 0;
    -ms-filter: "progid:DXImageTransform.Microsoft.Alpha(Opacity=0)";
    filter: alpha(opacity=0);
    -webkit-transform: translateY(-2000px) rotate(-45deg);
    -moz-transform: translateY(-2000px) rotate(-45deg);
    -o-transform: translateY(-2000px) rotate(-45deg);
    -ms-transform: translateY(-2000px) rotate(-45deg);
    transform: translateY(-2000px) rotate(-45deg);
  }
  60% {
    opacity: 1;
    -ms-filter: none;
    filter: none;
    -webkit-transform: translateY(30px) rotate(-45deg);
    -moz-transform: translateY(30px) rotate(-45deg);
    -o-transform: translateY(30px) rotate(-45deg);
    -ms-transform: translateY(30px) rotate(-45deg);
    transform: translateY(30px) rotate(-45deg);
  }
  80% {
    -webkit-transform: translateY(-10px) rotate(-45deg);
    -moz-transform: translateY(-10px) rotate(-45deg);
    -o-transform: translateY(-10px) rotate(-45deg);
    -ms-transform: translateY(-10px) rotate(-45deg);
    transform: translateY(-10px) rotate(-45deg);
  }
  100% {
    -webkit-transform: translateY(0) rotate(-45deg);
    -moz-transform: translateY(0) rotate(-45deg);
    -o-transform: translateY(0) rotate(-45deg);
    -ms-transform: translateY(0) rotate(-45deg);
    transform: translateY(0) rotate(-45deg);
  }
}
@-o-keyframes bounce {
  0% {
    opacity: 0;
    -ms-filter: "progid:DXImageTransform.Microsoft.Alpha(Opacity=0)";
    filter: alpha(opacity=0);
    -webkit-transform: translateY(-2000px) rotate(-45deg);
    -moz-transform: translateY(-2000px) rotate(-45deg);
    -o-transform: translateY(-2000px) rotate(-45deg);
    -ms-transform: translateY(-2000px) rotate(-45deg);
    transform: translateY(-2000px) rotate(-45deg);
  }
  60% {
    opacity: 1;
    -ms-filter: none;
    filter: none;
    -webkit-transform: translateY(30px) rotate(-45deg);
    -moz-transform: translateY(30px) rotate(-45deg);
    -o-transform: translateY(30px) rotate(-45deg);
    -ms-transform: translateY(30px) rotate(-45deg);
    transform: translateY(30px) rotate(-45deg);
  }
  80% {
    -webkit-transform: translateY(-10px) rotate(-45deg);
    -moz-transform: translateY(-10px) rotate(-45deg);
    -o-transform: translateY(-10px) rotate(-45deg);
    -ms-transform: translateY(-10px) rotate(-45deg);
    transform: translateY(-10px) rotate(-45deg);
  }
  100% {
    -webkit-transform: translateY(0) rotate(-45deg);
    -moz-transform: translateY(0) rotate(-45deg);
    -o-transform: translateY(0) rotate(-45deg);
    -ms-transform: translateY(0) rotate(-45deg);
    transform: translateY(0) rotate(-45deg);
  }
}
@keyframes bounce {
  0% {
    opacity: 0;
    -ms-filter: "progid:DXImageTransform.Microsoft.Alpha(Opacity=0)";
    filter: alpha(opacity=0);
    -webkit-transform: translateY(-2000px) rotate(-45deg);
    -moz-transform: translateY(-2000px) rotate(-45deg);
    -o-transform: translateY(-2000px) rotate(-45deg);
    -ms-transform: translateY(-2000px) rotate(-45deg);
    transform: translateY(-2000px) rotate(-45deg);
  }
  60% {
    opacity: 1;
    -ms-filter: none;
    filter: none;
    -webkit-transform: translateY(30px) rotate(-45deg);
    -moz-transform: translateY(30px) rotate(-45deg);
    -o-transform: translateY(30px) rotate(-45deg);
    -ms-transform: translateY(30px) rotate(-45deg);
    transform: translateY(30px) rotate(-45deg);
  }
  80% {
    -webkit-transform: translateY(-10px) rotate(-45deg);
    -moz-transform: translateY(-10px) rotate(-45deg);
    -o-transform: translateY(-10px) rotate(-45deg);
    -ms-transform: translateY(-10px) rotate(-45deg);
    transform: translateY(-10px) rotate(-45deg);
  }
  100% {
    -webkit-transform: translateY(0) rotate(-45deg);
    -moz-transform: translateY(0) rotate(-45deg);
    -o-transform: translateY(0) rotate(-45deg);
    -ms-transform: translateY(0) rotate(-45deg);
    transform: translateY(0) rotate(-45deg);
  }
}

</style>
