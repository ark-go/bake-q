<template>
  <img
    ref="refDrag"
    alt="Тындыр"
    src="/images/tandir-40.png"
    :style="{
      left: `${x}px`,
      top: `${y}px`,
      cursor: dragging ? 'grabbing' : 'grab',
    }"
    draggable="false"
    @mousedown="dragging = true"
    @touchmove="touchmove"
    @touchstart.stop
  />
</template>

<script>
import { ref, onMounted } from "vue";
import { useQuasar } from "quasar";
import { useIoSocket } from "src/stores/ioSocket";
export default {
  setup() {
    const ioSocket = useIoSocket();

    const refDrag = ref({});
    const $q = useQuasar();
    const dragging = ref(false);
    const mouseX = ref(0);
    const mouseY = ref(0);
    const x = ref(230);
    const y = ref(1);
    function listeningGmame(val) {
      x.value = val.x;
      y.value = val.y;
      if (refDrag.value.offsetLeft < 0) x.value = 0; // меньше левого
      if (refDrag.value.offsetTop < 0) y.value = 0; // меньше верха
      if (x.value + refDrag.value.offsetWidth > $q.screen.width)
        // больше экрана вправо
        x.value = $q.screen.width - refDrag.value.offsetWidth;
      if (y.value + refDrag.value.offsetHeight > $q.screen.height)
        // больше экрана вниз
        y.value = $q.screen.height - refDrag.value.offsetHeight;
    }
    ioSocket.socket.on("gameCursor", listeningGmame);

    onMounted(() => {
      console.dir(refDrag.value);
      window.addEventListener("mousemove", (e) => {
        if (dragging.value) {
          const diffX = e.clientX - mouseX.value;
          const diffY = e.clientY - mouseY.value;
          x.value += diffX;
          y.value += diffY;
          ioSocket.socket.emit("gameCursor", { x: x.value, y: y.value });
          if (refDrag.value.offsetLeft < 0) x.value = 0; // меньше левого
          if (refDrag.value.offsetTop < 0) y.value = 0; // меньше верха
          if (x.value + refDrag.value.offsetWidth > $q.screen.width)
            // больше экрана вправо
            x.value = $q.screen.width - refDrag.value.offsetWidth;
          if (y.value + refDrag.value.offsetHeight > $q.screen.height)
            // больше экрана вниз
            y.value = $q.screen.height - refDrag.value.offsetHeight;
        }
        mouseX.value = e.clientX;
        mouseY.value = e.clientY;
      });

      window.addEventListener("mouseup", () => {
        dragging.value = false;
      });
    });
    function touchmove(event) {
      // сразу ловим и движение и координаты для мобилок
      // if (dragging.value) {
      const diffX = event.touches[0].clientX - mouseX.value;
      const diffY = event.touches[0].clientY - mouseY.value;
      x.value += diffX;
      y.value += diffY;
      ioSocket.socket.emit("gameCursor", { x: x.value, y: y.value });
      if (refDrag.value.offsetLeft < 0) x.value = 0; // меньше левого
      if (refDrag.value.offsetTop < 0) y.value = 0; // меньше верха
      if (x.value + refDrag.value.offsetWidth > $q.screen.width) {
        // больше экрана вправо
        x.value = $q.screen.width - refDrag.value.offsetWidth;
        mouseX.value = x.value;
      } else {
        mouseX.value = event.touches[0].clientX;
      }
      if (y.value + refDrag.value.offsetHeight > $q.screen.height) {
        // больше экрана вниз
        y.value = $q.screen.height - refDrag.value.offsetHeight;
        mouseY.value = y.value;
      } else {
        mouseY.value = event.touches[0].clientY;
      }
      console.log(
        mouseY.value,
        event.touches[0].clientY,
        refDrag.value.offsetTop
      );
    }
    return {
      x,
      y,
      dragging,
      refDrag,
      touchmove,
    };
  },
};
</script>

<style scoped>
/* .drag-container {
  width: 100vw;
  height: 100vh;
} */

img {
  cursor: grab;
  /* position: relative; */
  position: fixed;
  z-index: 2000;
}
</style>
