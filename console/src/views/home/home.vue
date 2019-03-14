<template>
  <div>
    <Collapse>
      <Panel name="1">
        选项
        <div slot="content">
          <!-- <Row>
            <Col span="24">
              <Button type="primary" shape="circle" icon="ios-search" style="float: right"  @click="reset" ghost></Button>
            </Col>
          </Row> -->
        </div>
      </Panel>
    </Collapse>
    <br>
    <Row>
      <Col span="6">
        <Card  dis-hover>
          <p slot="title">
            请求数量
          </p>
          <div style="text-align:center; font-size:35px; font-weight:400">
            9999
          </div>
        </Card>
      </Col>
      <Col span="6">
        <Card  dis-hover>
          <p slot="title">
            请求数量
          </p>
          <div style="text-align:center; font-size:35px; font-weight:400">
            9999
          </div>
        </Card>
      </Col>
      <Col span="6">
        <Card  dis-hover>
          <p slot="title">
            请求数量
          </p>
          <div style="text-align:center; font-size:35px; font-weight:400">
            9999
          </div>
        </Card>
      </Col>
      <Col span="6">
        <Card  dis-hover>
          <p slot="title">
            请求数量
          </p>
          <div style="text-align:center; font-size:35px; font-weight:400">
            9999
          </div>
        </Card>
      </Col>
    </Row>
    <br>
    <Row>
      <Col span="12">
        <Card dis-hover>
          <p slot="title">
            请求数量
          </p>
          <ve-line :data="chartData" :settings="lineChartSettings"></ve-line>
        </Card>
      </Col>
      <Col span="12">
        <Card dis-hover>
          <p slot="title">
            请求数量
          </p>
          <ve-histogram :data="chartData" :settings="barChartSettings"></ve-histogram>
        </Card>
      </Col>
      <Col span="12">
        <Card dis-hover>
          <p slot="title">
            请求数量
          </p>
          <ve-bar :data="chartData" :settings="hBarChartSettings"></ve-bar>
        </Card>
      </Col>
      <Col span="12">
        <Card dis-hover>
          <p slot="title">
            请求数量
          </p>
          <ve-pie :data="chartData" :settings="circleChartSettings"></ve-pie>
        </Card>
      </Col>
    </Row>
  </div>
</template>

<script>
export default {
  data () {
    this.lineChartSettings = {
      stack: { '用户': ['访问用户', '下单用户'] },
      area: true
    };
    this.barChartSettings = {
      metrics: ['访问用户', '下单用户'],
      dimension: ['日期']
    };
    this.hBarChartSettings = {
      metrics: ['访问用户'],
      dataOrder: {
        label: '访问用户',
        order: 'desc'
      }
    };
    this.circleChartSettings = {
      roseType: 'radius'
    };
    return {
      chartData: {
        columns: ['日期', '访问用户', '下单用户', '下单率'],
        rows: [
          { '日期': '1/1', '访问用户': 1393, '下单用户': 1093, '下单率': 0.32 },
          { '日期': '1/2', '访问用户': 3530, '下单用户': 3230, '下单率': 0.26 },
          { '日期': '1/3', '访问用户': 2923, '下单用户': 2623, '下单率': 0.76 },
          { '日期': '1/4', '访问用户': 1723, '下单用户': 1423, '下单率': 0.49 },
          { '日期': '1/5', '访问用户': 3792, '下单用户': 3492, '下单率': 0.323 },
          { '日期': '1/6', '访问用户': 4593, '下单用户': 4293, '下单率': 0.78 }
        ]
      }
    };
  }
};
</script>
