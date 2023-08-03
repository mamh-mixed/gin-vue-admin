<template>
  <div class="cardList-wrapper">

    <div class="card-item">
      <div class="card-title">镜像</div>
      <div class="all-count">
        <span>5</span>
        <span class="suffix">个</span>
      </div>
    </div>

    <div class="card-item">
      <div class="card-title">已启动容器</div>
      <div class="all-count">
        <span>5</span>
        <span class="suffix">个</span>
      </div>
      <div class="card-info">
        <span>容器总数：<t id='containerCount'>20</t>个</span>
      </div>
    </div>

    <div class="card-item">
      <div class="card-title">本地数据卷</div>
      <div class="all-count">
        <span>5</span>
        <span class="suffix">个</span>
      </div>
      <div class="card-info">
        <span>总数据卷：<t id='volumeTotal'>0</t>个</span>
      </div>
    </div>

    <div class="card-item">
      <div class="card-title">集群节点</div>
      <div class="all-count">
        <span>5</span>
        <span class="suffix">个</span>
      </div>
      <div class="card-info">
        <span>管理结点：<t id='managerCount'>1</t></span>
      </div>
    </div>

  </div>
</template>

<script>

export default {
  name: 'CardInfo',
  props: ['title', 'className']

}
</script>

<style lang="scss" scoped>
.cardList-wrapper{
  display: flex;
  justify-content: space-between;
}

.cardList-wrapper .card-item:hover {
  -webkit-box-shadow: 3px 5px 15px rgb(40 40 40 / 36%);
  box-shadow: 3px 5px 15px rgb(40 40 40 / 36%);
  -webkit-transform: scale(1.03);
  transform: scale(1.03);
  -webkit-transition: all .1s ease-out;
  transition: all .1s ease-out;
}

.cardList-wrapper .card-item {
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  height: 149px;
  padding: 15px 15px 0;
  border-radius: 6px;
  cursor: pointer;
  margin: 0 5px 25px;
  width: 25%;
  color: #ffffff;
}

.cardList-wrapper .card-item:nth-child(1) {
  background: -webkit-gradient(linear, left top, right top, from(#5171fd), to(#c97afd));
  background: linear-gradient(90deg, #5171fd, #c97afd);
  -webkit-box-shadow: 0 5px 10px #c97afd;
  box-shadow: 0 5px 10px #c97afd;
}

.cardList-wrapper .card-item:nth-child(2) {
  background: -webkit-gradient(linear, left top, right top, from(#3dadf6), to(#737bfc));
  background: linear-gradient(90deg, #3dadf6, #737bfc);
  -webkit-box-shadow: 0 5px 10px #737bfc;
  box-shadow: 0 5px 10px #737bfc;
}

.cardList-wrapper .card-item:nth-child(3) {
  background: -webkit-gradient(linear, left top, right top, from(#ea677c), to(#ef9b5f));
  background: linear-gradient(90deg, #ea677c, #ef9b5f);
  -webkit-box-shadow: 0 5px 10px #ef9b5f;
  box-shadow: 0 5px 10px #ef9b5f;
}

.cardList-wrapper .card-item:nth-child(4) {
  background: -webkit-gradient(linear, left top, right top, from(#42d79b), to(#a6e25f));
  background: linear-gradient(90deg, #42d79b, #a6e25f);
  -webkit-box-shadow: 0 5px 10px #a6e25f;
  box-shadow: 0 5px 10px #a6e25f;
}

.suffix {
  font-size: .75rem;
}
.card-title {
  font-size: 1rem;
}
.all-count {
  font-size: 1.9rem;
  font-weight: 700;
  margin-top: 12px;
  padding-left: 25px;
}

.card-info {
  height: 46px;
  margin-top: 8px;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
}

.card-info span {
  font-size: .9rem;
  display: block;
}
</style>
