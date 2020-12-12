/*
 * @author: Haoyuan Liu
 * @date: 2020/12/2
 */

package util

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
