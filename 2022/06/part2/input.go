package main

var Input = `gzbzwzjwwrfftfrrbvvcbcvcffpssvhhzfzbfzzrbrtrcttnthnhsnhsnnclcpctpprwwgppppchhvvctvtsvswwcdcrrdzdhzdzvddvfvfjjlldvvfllfhlfhlhghnhghrgrsggvbbhdhffzggmttjqtjjgljjwtjwjggldlzzrfrpppjsjggqrrdbbgwgtwwwtqwtqqvccjnnhrhqqsnssjbssmjjbzjjqnqnlnjncjjzzqddfqqqnqvnqncqqmsshqqccssvpvllbbfvfmmmgzgdggtjgjtjbttngtnnhhfrfddmtmtftcftthvttwzwrzrbrgrbrqbqccmmtvtqvtqqzmzwmzzzsvzvhhhqlhqlqplpggvmvwmmshhnshsjszzljlnlrnllgddqllpmllsgswwqnqlnlqqlmqqvrqvrqrppvrrqwqhqllgjgcggsllqgllghlghhbvhhcvvbddvcvwwwzdzzfjjvcczbbwccjvccmvmfvfgfmfttwrttpmtmsszccjggcssmqsmsgszzfpphghvgvvcddtltwlwnnrbbfdbbtssrvvmwwjfflmlblwlrlmlwwhrhnhqnnvdvllzblzzndzdlzzhpprbrprnrzrjjlnlnrrjddgrrwrttcstsgshszzvgvfvpfpqfpfbpbfbwbmmzvvnnqvvqppzspzpzplldqlqsqlldmlmgmcmwmsmfsfzfdzzbpzzbbcwbbtppphlhshchjjszztctpprqppszzchcqczczddscsrstsntstjsjqjjfggvccgbggqhghrhppdldnlndnpddllgplgppnlpnnmwwglghhftfrtrqtqsszrssfjjzhzfhfssflfslsmmgrrcdrdnrdnnqznndrrbtrrmhrhbhsbsgsnslsqlqnncpccplpldpldppshsdssqdsqdsqdsdsdfsddchczhczcszcscffzhfzffpbbgttdhthnhnjnwnnwrrmwmlljzlzlzdzrrcvcbbsfstfsfnfqftqffgjfjzfjfsjswswccsvcvlccshhjpjcccnfnqqmsqmmrwmwjwswgsgjsjttmjmjttpqtqwttgmgfmmnbnrnpnwppgspggwnwqnqpqmmbmmnwnllbwwdqqzsqqtgttnrtrllcwcchpchcbhbpbnpbpsptpvtvzvrzzslljddjzjjmqmmdpmddcdzzrdrqddtctppwzzwrzrvzrvrzzjtztqqsmmcncllhmlmpmhgsrtrzjhjbtfvhmzpssfbjwcdshthnmmqmfhlhwcbbllwzbwfgfvzjcqblchzqqqgcgmpnbnnblhbcpgmvsbvtcmhsghldlhqlghgtpdvjflmjsmppdsvrjlwvhmwsmfvvdnzpmtfqjqpjdctnnrlfjfvdtmvdmhsczlfsqwfrqtlqwnzdzrcdzmhvrgwnjjtqjrqljhgcffglvhnsdssqdpfrfhtjwlqzvfjmpjnqhsjvndtstqjsqgcmgqdjvfqnlmtjlvcblndprmffrgqdnnncnlfflclgqbjbmsdqsjrmzpwtbqqqqsqmgthhjfqwzvcbqwlvdbffhcvncpldmchnptbnlbhmzrrjhzvzdrvtlhsnfnfdnvhlrlrmdcpnmvdwswctcqldszlvftqtwldrhmfjmfvcgjcdjsbjqjwdtslblwhqfvgrfcpnhszqqsfwbwcmvfvccvztdmcjqfjvdvmbdtcjvtwqpwsqjtdvpvvvsdrrvngmjztgjtnpdbmtlrbrjlwsdnthgzgpssgbzzrqvgblqhrtfbflnphvhpzmdfrwqvjjvcpsmdrqwdbzlpnqpmglgqzfhctrzdpzdthqgjtvpwcrlsmqnjwgzlnqbthvfswhjtrrsrswhbmnddgzmznvppbnmtjpzpdpmpzpmsgfstzldgmrtgplwwsbztphtcvdfgsqzqwrmlqpnhvpcqpfjpbrthrtwgqlfnrqcrpvhssjmhfgpnzzvlrlcbnpmddhtdvvfrrvprqrwbhfgvvltgrhrpwsdgvrlgthbztcgcfggtcfzqtlcdhpmcvpgcszslhpfrttnrdrqpqlgdfwtccmbhfrnlbhhmrtrjmzstbqhmtphzcpcllzsnghfvlwvzzvlqrjmfsvhrnjnvldgnbqvpjsmmphhrmhqrtcncmjwbdlqtrvmhgjrjsrddcpqnjhfmczfgwzspnrjfwvfdpdlcfpvctfrlspdwwlnpbvbglzsmgfsmrshsqgcvlfrvjssrglbwvvvgpvtqshbtqmzbnglflhhldtfzqssptrbnnzdqwqtstpftdqgmmhfjdlfwtmcmtmcgcvtfhzvsbllgbchvlhrgnvvbsnrwqrlvdqlcwwlgbjrvrzgcvljqzvdngtbhpnppjrzpmbwztzvnvrbfccfgvnqvrcmpdblngcvlwjwzpbbwmnslsdjmbrwbvnjsgcmsfvzvnwbzrrlzvgdnzcqqgggvmcwczjqnrddnzhlndgzjbvtbtjqdnlvlflqnfvjdmfstpdfqsgjdslgtpgdnvvpmfwvlqbmbgdrqjhdfhlmsdtfwpscsvdzmswszjfwqtsjqpfnjjlnsspvlgzwwtvwgdzfjjljfwbvfbvpglcqcdbdpshwcqzswbwhftbfqblzpqfmqpvzdjsrcqtvjntnhmlhmzllffcjsdnwpfzdglvlrhrljbrhjhgdnfcqcrccwrbbhrvqwrzlwjrrwzsfcwsvbqsjgtgpzqwlczljrtdhtzcgsfgqssdbjjmttdtzhrqtbqjtqfwmcpdftrfjmznpscsqmtfcdpcdjwcqjvmhngcqnmtmwfttcvpwgcnhhgnnbgjdvztzhczvqljjqwcmwcbvmqqjsfnmhtbtsvsnsfwfzgdvlctrgdvbjqpgfrrlsnppmvhfbbclprlvcssnvtsgftmlpqrpjzrrphrzltbtgfwjqqvbgqjdpdgqvzppfzcbhdbbjttcdzclsphtlzfvnfqgpmqvpzrqgpdnbmsrqgsnfdwpvndzmsllgmnrlhnnzldwshtrrsrsmdncwrcjnmcjlwbbfpzcwzvldtgvbcvhnbhgjgwjcvslrfcwbqlrqldvpcgnbwbzzncncftrgbrwchfrrtnpqsjbcpzvplnbqdgtvnnrcwwfvbdzdrsfspvtbhflhsbqmlsjvfmpvjbfvcrdmgrfqsmqgfrntfqnlqbvmsqtpncjrstspbqvfmddmhwsssdcddshmwdlscttmdzljtpwhzhzhwsdnmgjstfmlnqqvzzdqpqsjdsllswmqcjtnthwqnhbscrjdstljqgncjvjvlpfrtscrzrqghrdvdnbtnpshpldcchljzrzqjlwwscnphvrwlvzttcdjdrddgmvqpvdmttdqhwpfmslzvnrwlrrdttbhctgpgzjrmdwjbcmsprwggvmdmmltldgpbfnppnrpwcnpgtblccdvbsnfvgzmjppftvndmdslfshjvndfvvzjjlzhgfmhcggttrcrbrlwrqgjpchvhnjwqnbsnmftwszhzftglrfdvvnbcbzsslgmdchtrrrqqzhllrfhwfwbbdgfpdfwssmzcfcnzmrhfdddtdhfqmctsglqbwhwpnbdzbsfbbvvthrrgjvztjjwgjcgjntrddmmdldtmvjnwjbcqwfwvfhsrpchznhlqpcttqjffbrpbftftdjzlrqchmzrfgjrqlndfwfrghtsfsblcvtjmjrbfrwdgsgzmmjpdlbwzfscfsfcdqwdwnjwjbvvbptmfrqltmnlpbtrqspwdfmhnncqgtrtmbzhfmwrbcqhmmpmvprvwrjplspcmmspldmbgpbtmqjtrfcpfhcnpjbnlhjpzflstjqfqvzcnfgvrmtplndchffzrtfrqdpdnzrspddwmpzlfchrzzfcfvmbvfnlfwtfbvnffdqhljbvwwdtmszgrjtzwqdbgvvfphcnsdgvlmslqngfmsbrztrnpjprghmjffscbnfqwrvjjjtfzrmjtzbwdsmzgmbtjzvddhngmzvflwftblbzfd`
