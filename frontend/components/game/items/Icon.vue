<template>
  <i :class="`game-item-icon ${item}`" :title="item"></i>
</template>

<script>

export default {
  props: ["item"],
}
</script>


<style>
.game-item-icon {
  display: inline-block;
  width: 32px;
  height: 32px;
  background-image: url("~assets/img/icons/game-items-icons.png");
  &.backpack {
    background-position: 224px 608px;
  }
  &.axe {
    background-position: 480px 546px;
  }
  &.pickaxe {
    background-position: 448px 546px;
  }
  &.stone_hammer {
    background-position: 384px 546px;
  }
  &.stone_spear {
    background-position: 768px 675px;
  }
  &.stone_knife {
    background-position: 800px 707px;
  }
  &.log {
    background-position: 512px 320px;
  }
  &.stone {
    background-position: 480px 320px;
  }
  &.fire_dragon_egg {
    background-position: 320px 386px;
  }
}
</style>
