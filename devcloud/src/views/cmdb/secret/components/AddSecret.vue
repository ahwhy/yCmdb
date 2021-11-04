<template>
  <div>
    <!-- 添加凭证表单 -->
    <el-drawer
      title="添加凭证"
      :visible.sync="isVisible"
      :before-close="handleClose"
      size="40%"
    >
      <span>我来啦!</span>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: "AddSecret",
  props: {
    visible: {
      default: false,
      type: Boolean,
    },
  },
  methods: {
    handleClose() {
      this.$emit("update:visible", false);
    },
  },
  computed: {
    isVisible: {
      get() {
        return this.visible;
      },
      set(val) {
        this.$emit("update:visible", val);
      },
    },
  },
};
</script>