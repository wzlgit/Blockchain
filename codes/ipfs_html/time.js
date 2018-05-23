// 获取当前时间
function getCurrentTime() {
  var date = new Date();
  var week;
  switch (date.getDay()) {
    case 1:
      week = "星期一";
      break;
    case 2:
      week = "星期二";
      break;
    case 3:
      week = "星期三";
      break;
    case 4:
      week = "星期四";
      break;
    case 5:
      week = "星期五";
      break;
    case 6:
      week = "星期六";
      break;
    default:
      week = "星期天";
  }
  var years = date.getFullYear();
  var month = addZero(date.getMonth() + 1);
  var days = addZero(date.getDate());
  var hours = addZero(date.getHours());
  var minutes = addZero(date.getMinutes());
  var seconds = addZero(date.getSeconds());
  var formatDate = years + "年" + month + "月" + days + "日 " + hours + ":" +
    minutes + ":" + seconds + " " + week;
  return formatDate
}

// 时间前面补0，如12:05:09
function addZero(time) {
  if (time < 10) {
    return "0" + time;
  } else {
    return time;
  }
}

// 更新时间
function updateTime() {
  let dateDiv = document.getElementById("dateDiv");
  dateDiv.innerHTML = getCurrentTime();
}
