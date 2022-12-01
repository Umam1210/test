const input = "Saat meng*ecat tembok, Agung dib_antu oleh Raihan";
const input2 = "Berapa u(mur minimal[ untuk !mengurus ktp?";
const input3 = "Masing-masing anak mendap(atkan uang jajan ya=ng be&rbeda";

function containsSpecialChars(str) {
  const specialChars = "[`!@#$%^&*()_+=[]{};':\"\\|.<>/~]/";
  return specialChars
    .split("")
    .some((specialChar) => str.includes(specialChar));
}

const countSentenceWithoutSpecialChar = (input) => {
  const sentence = [];
  let data = input.split(" ");
  for (let i = 0; i < data.length; i++) {
    if (!containsSpecialChars(data[i])) {
      sentence.push(data[i]);
    }
  }
  return sentence.length;
};

console.log(countSentenceWithoutSpecialChar(input));
console.log(countSentenceWithoutSpecialChar(input2));
console.log(countSentenceWithoutSpecialChar(input3));