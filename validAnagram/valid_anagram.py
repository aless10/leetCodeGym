use std::collections::HashMap;

impl Solution {

    pub fn counter(s: String) -> HashMap<char, i32> {
        let mut s_map: HashMap<char, i32> = HashMap::new();
        for c in s.chars() {
            let x = s_map.entry(c).or_insert(0);
            *x += 1
        }
        return s_map;
    }

    pub fn is_anagram(s: String, t: String) -> bool {
        if (s.len() != t.len()) {
            return false;
        }

        let mut s_map: HashMap<char, i32> = Solution::counter(s);
        let mut t_map: HashMap<char, i32> = Solution::counter(t);

        for (key, val) in t_map.iter() {
            let s_value = s_map.get(&key);
            if (s_value != Some(val)) {
                return false;
            }
        }

        return true;
        
    }
}
