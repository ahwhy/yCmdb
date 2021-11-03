<template>
  <div class="main-container">
    <tips :tips="tips" />

    <div class="table-op">
      <div class="search">
        <el-input
          v-model="query.keywords"
          placeholder="请输入实例ID|名称|IP 敲回车进行搜索"
          @keyup.enter.native="get_hosts"
        ></el-input>
      </div>
      <div class="op"></div>
    </div>

    <div class="box-shadow">
      <el-table :data="hosts" v-loading="fetchHostLoading" style="width: 100%">
          <el-table-column prop="name" label="名称">
              <template slot-scope="{ row }">
                <router-link
                  :to="{ path: '/cmdb/host/detail', query: { id: row.id } }"
                >
                  {{ row.resource_id }}
                </router-link>
                <br />
                {{ row.name }}
              </template>
          </el-table-column>

          <el-table-column prop="name" label="资产来源">
              <template slot-scope="{ row }">
                {{ row.vendor }} <br />
                {{ row.region }}
              </template>
          </el-table-column>

          <el-table-column prop="name" label="内网IP/外网IP">
              <template slot-scope="{ row }">
                {{ row.private_ip }} <br />
                {{ row.public_ip }}
              </template>
          </el-table-column>

          <el-table-column prop="name" label="系统类型">
              <template slot-scope="{ row }">
                {{ row.os_name }}
              </template>
          </el-table-column>

          <el-table-column prop="sync_at" label="创建时间">
              <template slot-scope="scope">
                {{ scope.row.create_at | parseTime }}
              </template>
          </el-table-column>

          <el-table-column prop="expire_at" label="过期时间">
              <template slot-scope="scope">
                {{ scope.row.expire_at | parseTime }}
              </template>
          </el-table-column>

          <el-table-column prop="name" label="规格">
              <template slot-scope="{ row }">
                {{ row.cpu }} / {{ row.memory }}
              </template>
          </el-table-column>

          <el-table-column prop="name" label="状态">
              <template slot-scope="{ row }">
                {{ row.status }}
              </template>
          </el-table-column>
          
          <el-table-column prop="操作" align="center" label="状态">
            <template>
              <el-button type="text" disabled>归档</el-button>
              <el-button type="text" disabled>监控</el-button>
            </template>
          </el-table-column>
      </el-table>

      <pagination 
        v-show="total>0" 
        :total="total" 
        :page.sync="query.page_number" 
        :limit.sync="query.page_size" 
        @pagination="get_hosts" 
      />
    </div>
  </div>
</template>

<script>
import Tips from '@/components/Tips'
import Pagination from '@/components/Pagination'
import { LIST_HOST } from '@/api/cmdb/host.js'

const tips = [
  '现在仅同步了阿里云主机资产'
]

export default {
  name: 'CmdbHost',
  components: { Tips, Pagination },
  data() {
    return {
      tips: tips,
      fetchHostLoading: false,
      query: {page_size: 20, page_number: 1, keywords: "" },
      total: 0,
      hosts: []
    }
  },
  created() {
    this.get_hosts()
  },
  methods: {
    async get_hosts() {
      this.fetchHostLoading = true;
      try {
        const resp = await LIST_HOST(this.query);
        this.hosts = resp.data.items;
        this.total = resp.data.total;
      } catch (error) {
        this.$notify.error({
          title: "获取主机异常",
          message: error,
        });
      } finally {
        this.fetchHostLoading = false;
      }
    },
    handleSizeChange(val) {
      this.query.page_size = val
      this.get_hosts()
    },
    handleCurrentChange(val) {
      this.query.page_number = val
      this.get_hosts()
    }
  }
}
</script>

<style lang="scss" scoped>
.box-shadow {
  margin: 12px 0;
}

.table-op {
  margin-top: 12px;
}

.search {
  width: 30%;
}
</style>