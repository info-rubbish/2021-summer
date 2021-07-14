<template>
  <div
    class="
      w-screen
      h-screen
      fixed
      bg-gray-700 bg-opacity-70
      backdrop-filter backdrop-blur
      flex
      items-center
      justify-center
      z-50
    "
  >
    <div
      class="p-10 rounded-lg bg-white border-4 text-center"
      :class="[levelBorderColor]"
    >
      <p><slot></slot></p>
      <button
        autofocus
        class="mt-6 p-4 rounded"
        :class="[levelBgColor]"
        @click="$emit('close')"
      >
        OK
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: ["level"],
  emits: ["close"],
  name: "alert",
  setup(props) {
    var bg, border;
    switch (props.level) {
      case 1:
        bg = "bg-yellow-400";
        border = "border-yellow-400";
        break;
      case 2:
        bg = "bg-red-600";
        border = "border-red-600";
        break;
      default:
        bg = "bg-blue-500";
        border = "border-blue-500";
        break;
    }
    return { levelBgColor: bg, levelBorderColor: border };
  },
};
</script>