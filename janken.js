let rock = document.getElementById("rock");
let scissors = document.getElementById("scissors");
let paper = document.getElementById("paper");
let result = document.getElementById("result");
let yourhand = document.getElementById("cpu");
let win = document.getElementById("win");
let lose = document.getElementById("lose");
let draw = document.getElementById("draw");
let wincount = 0;
let losecount = 0;
let drawcount = 0;
rock.addEventListener("click",function(){
const cpuhand = Math.floor(Math.random() * 3);
const myhand = 0;
CpuHand(cpuhand);
VictoryLossJudgment(myhand,cpuhand);
});
scissors.addEventListener("click",function(){
const cpuhand = Math.floor(Math.random() * 3) ;
const myhand = 1;
CpuHand(cpuhand);
VictoryLossJudgment(myhand,cpuhand);
});
paper.addEventListener("click",function(){
const cpuhand = Math.floor(Math.random() * 3);
const myhand = 2;
CpuHand(cpuhand);
VictoryLossJudgment(myhand,cpuhand);
});

function CpuHand(hand){
  const hands　= ["グー","チョキ","パー"];
  yourhand.innerText = "相手の手："+ hands[hand];
}
function VictoryLossJudgment(my,cpu){
  if (my === cpu){
    result.innerText = "結果：あいこ";
    drawcount++;
    draw.innerText = drawcount;
  } else if(
    (my === 0 && cpu === 1) ||
    (my === 1 && cpu === 2) ||
    (my === 2 && cpu === 0)
    ){
      result.innerText = "結果：勝ち";
      wincount++;
      win.innerText = wincount;
  }else{
      result.innerText = "結果：負け";
      losecount++;
      lose.innerText = losecount;

  }  
}


