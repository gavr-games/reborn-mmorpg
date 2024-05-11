<template>
  <i :class="`game-item-icon ${item}`" :title="item" />
</template>

<script>

export default {
  props: {
    item: {
      type: String,
      default: ''
    }
  }
}
</script>

<style lang="scss">
.game-item-icon {
  display: inline-block;
  width: 32px;
  height: 32px;
  background-image: url("~assets/img/icons/game-items-icons.png");
  &.animal_skin {
    background-position: 256px 320px;
  }
  &.axe {
    background-position: 480px 546px;
  }
  &.backpack {
    background-position: 224px 608px;
  }
  &.bone {
    background-position: 224px 320px;
  }
  &.bone_needle {
    background-position: 320px 514px;
  }
  &.bone_saw {
    background-position: 736px 739px;
  }
  &.cactus_slice {
    background-position: 448px 288px;
  }
  &.carrot {
    background-position: 320px 418px;
  }
  &.carrot_seed {
    background-position: 416px 482px;
  }
  &.claim_stone {
    background-position: 512px 514px;
  }
  &.dungeon_key {
    background-position: 224px 514px;
  }
  &.fire_dragon_egg {
    background-position: 320px 386px;
  }
  &.grass {
    background-position: 32px 514px;
  }
  &.gold {
    background-position: 192px 482px;
  }
  &.gold_ingot {
    background-position: 418px 320px;
  }
  &.golden_armor {
    background-position: 192px 608px;
  }
  &.healing_balm {
    background-position: 448px 578px;
  }
  &.iron_ingot {
    background-position: 160px 320px;
  }
  &.iron_nails {
    background-position: 800px 867px;
  }
  &.leather_robe {
    background-position: 224px 640px;
  }
  &.log {
    background-position: 512px 320px;
  }
  &.pickaxe {
    background-position: 448px 546px;
  }
  &.rope {
    background-position: 96px 546px;
  }
  &.small_bag {
    background-position: 320px 482px;
  }
  &.stone {
    background-position: 480px 320px;
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
  &.wooden_fishing_rod {
    background-position: 512px 352px;
  }
  &.wooden_shovel {
    background-position: 416px 546px;
  }
}
</style>
