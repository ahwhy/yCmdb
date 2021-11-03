<template>
  <div>
    <!-- 基础信息 -->
    <div class="box-shadow basic-info">
      <el-descriptions title="主机信息">
        <el-descriptions-item label="名称">
          {{ host.name }}
        </el-descriptions-item>
        <el-descriptions-item label="实例ID">
          {{ host.instance_id }}
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          {{ host.status }}
        </el-descriptions-item>
        <el-descriptions-item label="规格">
          {{ host.cpu }} / {{ host.memory }}
        </el-descriptions-item>
        <el-descriptions-item label="厂商">
          {{ host.vendor }}
        </el-descriptions-item>
        <el-descriptions-item label="系统">
          {{ host.os_name }}
        </el-descriptions-item>
        <el-descriptions-item label="地域">
          {{ host.region }}
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">
          {{ host.create_at | parseTime }}
        </el-descriptions-item>
        <el-descriptions-item label="序列号">
          {{ host.serial_number }}
        </el-descriptions-item>
        <el-descriptions-item label="过期时间">
          {{ host.expire_at | parseTime }}
        </el-descriptions-item>
        <el-descriptions-item label="同步时间">
          {{ host.sync_at | parseTime }}
        </el-descriptions-item>
        <el-descriptions-item label="同步账号">
          {{ host.sync_account }}
        </el-descriptions-item>
        <el-descriptions-item label="内网IP">
          {{ host.private_ip.join(",") }}
        </el-descriptions-item>
        <el-descriptions-item label="公网IP">
          {{ host.public_ip.join(",") }}
        </el-descriptions-item>
      </el-descriptions>
    </div>
    <!-- 关联信息 -->
    <el-card class="box-shadow associate-info">
      <el-tabs v-model="activeName">
        <el-tab-pane label="主机事件" name="event"> 主机变更 </el-tab-pane>
      </el-tabs>
    </el-card>
  </div>
</template>
<script>
import { GET_HOST } from "@/api/cmdb/host";

export default {
  name: "HostDetail",
  data() {
    return {
      host: {},
      activeName: "event",
    };
  },
  async mounted() {
    try {
      let resp = await GET_HOST(this.$route.query.id);
      this.host = resp.data;
    } catch (error) {
      this.$notify.error({
        title: "获取主机异常",
        message: error,
      });
    }
  },
};
</script>

<style scoped>
.basic-info {
  padding: 8px;
  background-color: white;
}
.associate-info {
  margin-top: 12px;
}
</style>
